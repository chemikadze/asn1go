package asn1go

import (
	"errors"
	"fmt"
	goast "go/ast"
	goprint "go/printer"
	gotoken "go/token"
	"io"
	"strings"
)

type CodeGenerator interface {
	Generate(module ModuleDefinition, writer io.Writer) error
}

type GenParams struct {
	Package string
	Prefix  string
	Type    GenType
}

type GenType int

const (
	GEN_DECLARATIONS GenType = iota
)

func NewCodeGenerator(params GenParams) CodeGenerator {
	switch params.Type {
	case GEN_DECLARATIONS:
		return &declCodeGen{params}
	default:
		return nil
	}
}

type declCodeGen struct {
	Params GenParams
}

type moduleContext struct {
	extensibilityImplied bool
	tagDefault           int
	errors               []error
	lookupContext        ModuleBody
	requiredModules      []string
}

func (ctx *moduleContext) appendError(err error) {
	ctx.errors = append(ctx.errors, err)
}

func (ctx *moduleContext) requireModule(module string) {
	for _, existing := range ctx.requiredModules {
		if existing == module {
			return
		}
	}
	ctx.requiredModules = append(ctx.requiredModules, module)
}

// Generate declarations from module.
//
// Feature support status:
// - [x] ModuleIdentifier
// - [ ] TagDefault
// - [ ] ExtensibilityImplied
// - [.] ModuleBody -- see generateDeclarations
func (gen declCodeGen) Generate(module ModuleDefinition, writer io.Writer) error {
	ctx := moduleContext{
		extensibilityImplied: module.ExtensibilityImplied,
		tagDefault:           module.TagDefault,
		lookupContext:        module.ModuleBody,
	}
	moduleName := goast.NewIdent(goifyName(module.ModuleIdentifier.Reference))
	if len(gen.Params.Package) > 0 {
		moduleName = goast.NewIdent(gen.Params.Package)
	}
	ast := &goast.File{
		Name:  moduleName,
		Decls: ctx.generateDeclarations(module),
	}
	if len(ctx.errors) != 0 {
		msg := "errors generating Go AST from module: \n"
		for _, err := range ctx.errors {
			msg += "  " + err.Error() + "\n"
		}
		return errors.New(msg)
	}
	importDecls := make([]goast.Decl, 0)
	for _, moduleName := range ctx.requiredModules {
		modulePath := &goast.BasicLit{Kind: gotoken.STRING, Value: fmt.Sprintf("\"%v\"", moduleName)}
		specs := []goast.Spec{&goast.ImportSpec{Path: modulePath}}
		importDecls = append(importDecls, &goast.GenDecl{Tok: gotoken.IMPORT, Specs: specs})
	}
	ast.Decls = append(importDecls, ast.Decls...)
	return goprint.Fprint(writer, gotoken.NewFileSet(), ast)
}

func goifyName(name string) string {
	return strings.Title(strings.Replace(name, "-", "_", -1))
}

// generateDeclarations based on ModuleBody of module
//
// Feature support status:
// - [.] AssignmentList
//    - [ ] ValueAssignment
//    - [x] TypeAssignment
// - [ ] Imports
func (ctx *moduleContext) generateDeclarations(module ModuleDefinition) []goast.Decl {
	decls := make([]goast.Decl, 0)
	for _, assignment := range module.ModuleBody.AssignmentList {
		switch a := assignment.(type) {
		case TypeAssignment:
			decls = append(decls, ctx.generateTypeDecl(a.TypeReference, a.Type))
		}
	}
	return decls
}

func (ctx *moduleContext) generateTypeDecl(reference TypeReference, typeDescr Type) goast.Decl {
	return &goast.GenDecl{
		Tok: gotoken.TYPE,
		Specs: []goast.Spec{
			&goast.TypeSpec{
				Name: goast.NewIdent(goifyName(reference.Name())),
				Type: ctx.generateTypeBody(typeDescr),
			},
		},
	}
}

func (ctx *moduleContext) generateTypeBody(typeDescr Type) goast.Expr {
	switch t := typeDescr.(type) {
	case BooleanType:
		return goast.NewIdent("bool")
	case IntegerType:
		return goast.NewIdent("int64") // TODO signed, unsigned, range constraints
	case CharacterStringType:
		return goast.NewIdent("string")
	case RealType:
		return goast.NewIdent("float64")
	case OctetStringType:
		return &goast.ArrayType{Elt: goast.NewIdent("byte")}
	case SequenceType:
		return ctx.structFromComponents(t.Components)
	case SetType:
		return ctx.structFromComponents(t.Components)
	case SequenceOfType:
		return &goast.ArrayType{Elt: ctx.generateTypeBody(t.Type)}
	case SetOfType:
		return &goast.ArrayType{Elt: ctx.generateTypeBody(t.Type)}
	case TaggedType: // TODO should put tags in go code?
		return ctx.generateTypeBody(t.Type)
	case ConstraintedType: // TODO should generate checking code?
		return ctx.generateTypeBody(t.Type)
	case TypeReference: // TODO should useful types be separate type by itself?
		nameAndType := ctx.resolveTypeReference(t)
		if nameAndType != nil {
			specialCase := ctx.generateSpecialCase(*nameAndType)
			if specialCase != nil {
				return specialCase
			}
		}
		return goast.NewIdent(goifyName(t.Name()))
	case RestrictedStringType: // TODO should generate checking code?
		return goast.NewIdent("string")
	case BitStringType:
		ctx.requireModule("encoding/asn1")
		return goast.NewIdent("asn1.BitString")
	case EnumeratedType:
		ctx.requireModule("encoding/asn1")
		return goast.NewIdent("asn1.Enumerated")
	case AnyType:
		return &goast.InterfaceType{}
	case ObjectIdentifierType:
		ctx.requireModule("encoding/asn1")
		return goast.NewIdent("asn1.ObjectIdentifier")
	default:
		// NullType
		// ChoiceType
		// RestrictedStringType
		ctx.appendError(fmt.Errorf("ignoring unsupported type %#v", typeDescr))
		return nil
	}
}

func (ctx *moduleContext) structFromComponents(components ComponentTypeList) goast.Expr {
	fields := &goast.FieldList{}
	for _, field := range components {
		switch f := field.(type) {
		case NamedComponentType:
			fields.List = append(fields.List, ctx.generateStructField(f))
		case ComponentsOfComponentType: // TODO
		}
	}
	return &goast.StructType{
		Fields: fields,
	}
}

func (ctx *moduleContext) generateStructField(f NamedComponentType) *goast.Field {
	return &goast.Field{
		Names: append(make([]*goast.Ident, 0), goast.NewIdent(goifyName(f.NamedType.Identifier.Name()))),
		Type:  ctx.generateTypeBody(f.NamedType.Type),
		Tag:   ctx.asn1TagFromType(f),
	}
}

func (ctx *moduleContext) asn1TagFromType(nt NamedComponentType) *goast.BasicLit {
	t := nt.NamedType.Type
	components := make([]string, 0)
	if nt.IsOptional {
		components = append(components, "optional")
	}
	if nt.Default != nil {
		if defaultNumber, ok := (*nt.Default).(Number); ok {
			components = append(components, fmt.Sprintf("default:%v", defaultNumber.IntValue()))
		}
	}
	// unwrap type
unwrap:
	for {
		switch tt := t.(type) {
		case TaggedType:
			switch tt.Tag.Class {
			case CLASS_APPLICATION:
				components = append(components, "application")
			case CLASS_PRIVATE:
				components = append(components, "private")
			}
			if tt.TagType == TAGS_EXPLICIT {
				components = append(components, "explicit")
			}
			switch cn := ctx.lookupValue(tt.Tag.ClassNumber).(type) {
			case Number:
				components = append(components, fmt.Sprintf("tag:%v", cn.IntValue()))
			default:
				ctx.appendError(fmt.Errorf("tag value should be Number, got %#v", cn))
			}
			t = tt.Type
		case ConstraintedType:
			t = tt.Type
		default:
			break unwrap
		}
	}
	// add type-specific tags
	switch tt := t.(type) {
	case RestrictedStringType:
		switch tt.LexType {
		case IA5String:
			components = append(components, "ia5")
		case UTF8String:
			components = append(components, "utf8")
		case PrintableString:
			components = append(components, "printable")
		}
	case SetType:
		components = append(components, "set")
	case SetOfType:
		components = append(components, "set")
	case TypeReference:
		switch ctx.unwrapToLeafType(tt).TypeReference.Name() {
		case GeneralizedTimeName:
			components = append(components, "generalized")
		case UTCTimeName:
			components = append(components, "utc")
		}
		// TODO set          causes a SET, rather than a SEQUENCE type to be expected
		// TODO omitempty    causes empty slices to be skipped\
	}
	if len(components) > 0 {
		return &goast.BasicLit{
			Value: fmt.Sprintf("`asn1:\"%s\"`", strings.Join(components, ",")),
			Kind:  gotoken.STRING,
		}
	} else {
		return nil
	}
}

func (ctx *moduleContext) generateSpecialCase(resolved TypeAssignment) goast.Expr {
	if resolved.TypeReference.Name() == GeneralizedTimeName || resolved.TypeReference.Name() == UTCTimeName {
		// time types in encoding/asn1go don't support wrapping of time.Time
		ctx.requireModule("time")
		return goast.NewIdent("time.Time")
	} else if _, ok := ctx.removeWrapperTypes(resolved.Type).(BitStringType); ok {
		ctx.requireModule("encoding/asn1")
		return goast.NewIdent("asn1.BitString")
	}
	return nil
}

// TODO really lookup values from module and imports
func (ctx *moduleContext) lookupValue(val Value) Value {
	return val
}

// resolveTypeReference resolves references until reaches unresolved type, useful type, or declared type
// returns type reference of most nested type which is not type reference itself
// returns nil if type is not resolved
func (ctx *moduleContext) resolveTypeReference(reference TypeReference) *TypeAssignment {
	unwrapped := ctx.unwrapToLeafType(reference)
	if unwrapped.Type != nil {
		return &unwrapped
	} else if tt := ctx.lookupUsefulType(unwrapped.TypeReference); tt != nil {
		return &TypeAssignment{unwrapped.TypeReference, tt}
	} else {
		ctx.appendError(fmt.Errorf("can not resolve TypeReference %v", reference.Name()))
		return nil
	}
}

func (ctx *moduleContext) lookupUsefulType(reference TypeReference) Type {
	if usefulType, ok := USEFUL_TYPES[reference.Name()]; ok {
		return usefulType
	} else {
		return nil
	}
}

func (ctx *moduleContext) removeWrapperTypes(t Type) Type {
	for {
		switch tt := t.(type) {
		case TaggedType:
			t = tt.Type
		case ConstraintedType:
			t = tt.Type
		default:
			return t
		}
	}
}

// unwrapToLeafType walks over transitive type references, tags and constraints and yields "root" type reference
func (ctx *moduleContext) unwrapToLeafType(reference TypeReference) TypeAssignment {
	if assignment := ctx.lookupContext.AssignmentList.GetType(reference.Name()); assignment != nil {
		t := assignment.Type
		if tt, ok := ctx.removeWrapperTypes(t).(TypeReference); ok {
			return ctx.unwrapToLeafType(tt)
		} else {
			return *assignment
		}
	}
	return TypeAssignment{reference, nil}
}
