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

// CodeGenerator is an interface for code generation from ASN.1 modules.
type CodeGenerator interface {
	Generate(module ModuleDefinition, writer io.Writer) error
}

// GenParams is code generator configuration.
type GenParams struct {
	// Package is go package name.
	// If not specified, ASN.1 module name will be used to derive go module name.
	Package string
	// Type is a type of code generation to run.
	// TODO: deprecate in favor of separate New methods.
	Type GenType
	// IntegerRepr controls how INTEGER type is expressed in generated go code.
	IntegerRepr IntegerRepr
}

// GenType is code generator type.
type GenType int

const (
	// GEN_DECLARATIONS is code generator that is
	GEN_DECLARATIONS GenType = iota
)

// IntegerRepr is enum controlling how INTEGER is represented.
type IntegerRepr string

// IntegerRepr modes supported.
const (
	IntegerReprInt64  IntegerRepr = "int64"
	IntegerReprBigInt IntegerRepr = "big.Int"
)

// NewCodeGenerator creates a new code generator from provided params.
func NewCodeGenerator(params GenParams) CodeGenerator {
	if params.IntegerRepr == "" {
		params.IntegerRepr = IntegerReprInt64
	}
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

// moduleContext is context used to track state of the code generation.
type moduleContext struct {
	extensibilityImplied bool
	// tagDefault is a ModuleDefinition.TagDefault value.
	tagDefault int
	// errors collected during conversion.
	// TODO: switch to explicit error passing.
	errors        []error
	lookupContext ModuleBody
	// requiredModules holds go modules required by generated code.
	requiredModules []string
	params          GenParams
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

// Generate declarations from module to be used together with encoding/asn1.
//
// Feature support status:
// - [x] ModuleIdentifier
// - [x] TagDefault (except AUTOMATIC)
// - [ ] ExtensibilityImplied
// - [.] ModuleBody -- see moduleContext.generateDeclarations.
func (gen declCodeGen) Generate(module ModuleDefinition, writer io.Writer) error {
	if module.TagDefault == TAGS_AUTOMATIC {
		// See x.680, section 12.3. It implies certain transformations to component and alternative lists that are not implemented.
		return errors.New("AUTOMATIC tagged modules are not supported")
	}
	ctx := moduleContext{
		extensibilityImplied: module.ExtensibilityImplied,
		tagDefault:           module.TagDefault,
		lookupContext:        module.ModuleBody,
		params:               gen.Params,
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

// generateDeclarations produces go declarations based on ModuleBody of module.
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
	var isSet bool
	typeBody := ctx.generateTypeBody(typeDescr, &isSet)
	spec := &goast.TypeSpec{
		Name:   goast.NewIdent(goifyName(reference.Name())),
		Type:   typeBody,
		Assign: 1, // not a valid Pos, but formatter just needs non-empty value
	}
	decl := &goast.GenDecl{
		Tok:   gotoken.TYPE,
		Specs: []goast.Spec{spec},
	}
	if _, ok := typeBody.(*goast.StructType); ok {
		spec.Assign = 0
	}
	if isSet {
		oldName := spec.Name.Name
		spec.Name.Name += "SET"
		spec.Assign = 0
		newName := spec.Name.Name
		decl.Specs = append(decl.Specs, &goast.TypeSpec{
			Name:   goast.NewIdent(oldName),
			Assign: 1,
			Type:   goast.NewIdent(newName),
		})
	}
	return decl
}

func (ctx *moduleContext) generateTypeBody(typeDescr Type, isSet *bool) goast.Expr {
	switch t := typeDescr.(type) {
	case BooleanType:
		return goast.NewIdent("bool")
	case IntegerType:
		// TODO: generate consts
		switch ctx.params.IntegerRepr {
		case IntegerReprInt64:
			return goast.NewIdent("int64") // TODO signed, unsigned, range constraints
		case IntegerReprBigInt:
			ctx.requireModule("math/big")
			return &goast.StarExpr{X: goast.NewIdent("big.Int")}
		default:
			ctx.appendError(fmt.Errorf("unknown int type mode: %v", ctx.params.IntegerRepr))
			return goast.NewIdent("int64")
		}
	case CharacterStringType:
		return goast.NewIdent("string")
	case RealType:
		return goast.NewIdent("float64")
	case OctetStringType:
		return &goast.ArrayType{Elt: goast.NewIdent("byte")}
	case SequenceType:
		return ctx.structFromComponents(t.Components, t.ExtensionAdditions)
	case SetType:
		*isSet = true
		return ctx.structFromComponents(t.Components, t.ExtensionAdditions)
	case SequenceOfType:
		return &goast.ArrayType{Elt: ctx.generateTypeBody(t.Type, isSet)}
	case SetOfType:
		*isSet = true
		return &goast.ArrayType{Elt: ctx.generateTypeBody(t.Type, isSet)}
	case TaggedType: // TODO should put tags in go code?
		return ctx.generateTypeBody(t.Type, isSet)
	case ConstraintedType: // TODO should generate checking code?
		return ctx.generateTypeBody(t.Type, isSet)
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
		// TODO: generate consts
		ctx.requireModule("encoding/asn1")
		return goast.NewIdent("asn1.Enumerated")
	case AnyType:
		return &goast.InterfaceType{Methods: &goast.FieldList{}}
	case ObjectIdentifierType:
		ctx.requireModule("encoding/asn1")
		return goast.NewIdent("asn1.ObjectIdentifier")
	case ChoiceType:
		return ctx.generateChoiceType(t, isSet)
	default:
		// NullType
		ctx.appendError(fmt.Errorf("ignoring unsupported type %#v", typeDescr))
		return nil
	}
}

func (ctx *moduleContext) generateChoiceType(t ChoiceType, isSet *bool) goast.Expr {
	if ctx.hasTaggedAlternatives(t) {
		return goast.NewIdent("asn1.RawValue")
	}
	if len(t.AlternativeTypeList) == 1 {
		return ctx.generateTypeBody(t.AlternativeTypeList[0].Type, isSet) // optimization for X.509 edge case
	}
	return &goast.InterfaceType{Methods: &goast.FieldList{}}
}

func (ctx *moduleContext) hasTaggedAlternatives(t ChoiceType) bool {
	for _, f := range t.AlternativeTypeList {
		if ctx.taggedChoiceTypeAlternative(f.Identifier, f.Type) {
			return true
		}
	}
	return false
}

func (ctx *moduleContext) taggedChoiceTypeAlternative(name Identifier, t Type) bool {
	switch t := t.(type) {
	case TaggedType:
		return true
	case TypeReference:
		if t.Name() == GeneralizedTimeName || t.Name() == UTCTimeName {
			return false
		}
		realType := ctx.resolveTypeReference(t)
		if realType == nil {
			return false
		}
		return ctx.taggedChoiceTypeAlternative(name, realType.Type)
	case ConstraintedType:
		return ctx.taggedChoiceTypeAlternative(name, t.Type)
	default:
		return false
	}
}

func (ctx *moduleContext) structFromComponents(components ComponentTypeList, extensions ExtensionAdditions) goast.Expr {
	fields := &goast.FieldList{}
	for _, field := range components {
		switch f := field.(type) {
		case NamedComponentType:
			fields.List = append(fields.List, ctx.generateStructField(f))
		case ComponentsOfComponentType: // TODO: implement
			ctx.appendError(errors.New("COMPONENTS OF is not supported"))
		}
	}
	for _, field := range extensions {
		switch f := field.(type) {
		case NamedComponentType:
			fields.List = append(fields.List, ctx.generateStructField(f))
		case ComponentsOfComponentType: // TODO: implement
			ctx.appendError(errors.New("COMPONENTS OF is not supported"))
		}
	}
	return &goast.StructType{
		Fields: fields,
	}
}

func (ctx *moduleContext) generateStructField(f NamedComponentType) *goast.Field {
	var stubBool bool // we care about isSet / shouldAssign only for top-level decls
	return &goast.Field{
		Names: append(make([]*goast.Ident, 0), goast.NewIdent(goifyName(f.NamedType.Identifier.Name()))),
		Type:  ctx.generateTypeBody(f.NamedType.Type, &stubBool),
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
		if !nt.IsOptional { // ensure it's marked as optional
			components = append(components, "optional")
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
			tagType := ctx.tagDefault
			if tt.HasTagType {
				tagType = tt.TagType
			}
			switch tagType {
			case TAGS_EXPLICIT:
				components = append(components, "explicit")
			case TAGS_IMPLICIT: // nothing to do
			case TAGS_AUTOMATIC:
				ctx.appendError(fmt.Errorf("type %v: AUTOMATIC tags are not supported", nt.NamedType.Identifier))
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
		case NumericString:
			components = append(components, "numeric")
		case PrintableString:
			// default type
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
