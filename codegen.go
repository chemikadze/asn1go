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

func NewCodeGenerator(gentype int) CodeGenerator {
	switch gentype {
	case GEN_DECLARATIONS:
		return &declCodeGen{}
	default:
		return nil
	}
}

const (
	GEN_DECLARATIONS = iota
)

var (
	USEFUL_TYPES map[string]Type
)

type declCodeGen struct{}

type moduleContext struct {
	extensibilityImplied bool
	tagDefault           int
	errors               []error
	lookupContext        ModuleBody
}

func (ctx *moduleContext) appendError(err error) {
	ctx.errors = append(ctx.errors, err)
}

/** Generate declarations from module

Feature support status:
 - [x] ModuleIdentifier
 - [ ] TagDefault
 - [ ] ExtensibilityImplied
 - [.] ModuleBody -- see generateDeclarations
*/
func (gen declCodeGen) Generate(module ModuleDefinition, writer io.Writer) error {
	ctx := moduleContext{
		extensibilityImplied: module.ExtensibilityImplied,
		tagDefault:           module.TagDefault,
		lookupContext:        module.ModuleBody,
	}
	ast := &goast.File{
		Name:  goast.NewIdent(goifyName(module.ModuleIdentifier.Reference)),
		Decls: ctx.generateDeclarations(module),
	}
	if len(ctx.errors) != 0 {
		msg := "Errors generating Go AST from module: \n"
		for _, err := range ctx.errors {
			msg += "  " + err.Error() + "\n"
		}
		return errors.New(msg)
	}
	return goprint.Fprint(writer, gotoken.NewFileSet(), ast)
}

func goifyName(name string) string {
	return strings.Replace(name, "-", "_", -1)
}

/** generateDeclarations based on ModuleBody of module

Feature support status:
 - [.] AssignmentList
   - [ ] ValueAssignment
   - [x] TypeAssignment
 - [ ] Imports
*/
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
		fields := &goast.FieldList{}
		for _, field := range t.Components {
			switch f := field.(type) {
			case NamedComponentType:
				fields.List = append(fields.List, ctx.generateStructField(f))
			case ComponentsOfComponentType: // TODO
			}
		}
		return &goast.StructType{
			Fields: fields,
		}
	case SequenceOfType:
		return &goast.ArrayType{Elt: ctx.generateTypeBody(t.Type)}
	case TaggedType: // TODO should put tags in go code?
		return ctx.generateTypeBody(t.Type)
	case ConstraintedType: // TODO should generate checking code?
		return ctx.generateTypeBody(t.Type)
	case TypeReference: // TODO generate references instead of embedding
		return ctx.generateTypeBody(ctx.lookupType(t))
	case RestrictedStringType: // TODO should generate checking code?
		return goast.NewIdent("string")
	case BitStringType: // TODO
		return &goast.ArrayType{Elt: goast.NewIdent("bool")}
	default:
		// NullType
		// ObjectIdentifierType
		// ChoiceType
		// RestrictedStringType
		ctx.appendError(errors.New(fmt.Sprintf("Ignoring unsupported type %#v", typeDescr)))
		return nil
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
	/*
		TODO:
		set          causes a SET, rather than a SEQUENCE type to be expected
		omitempty:   causes empty slices to be skipped
		utc:         causes time.Time to be marshaled as ASN.1, UTCTime values
		generalized: causes time.Time to be marshaled as ASN.1, GeneralizedTime values
	*/
	switch t := t.(type) {
	case TaggedType:
		if t.Tag.Class == CLASS_APPLICATION {
			components = append(components, "application")
		}
		if t.TagType == TAGS_EXPLICIT {
			components = append(components, "explicit")
		}
		switch cn := ctx.lookupValue(t.Tag.ClassNumber).(type) {
		case Number:
			components = append(components, fmt.Sprintf("tag:%v", cn.IntValue()))
		default:
			ctx.appendError(errors.New(fmt.Sprintf("Tag value should be Number, got %#v", cn)))
		}
	case RestrictedStringType:
		switch t.LexType {
		case IA5String:
			components = append(components, "ia5")
		case UTF8String:
			components = append(components, "utf8")
		case PrintableString:
			components = append(components, "printable")
		}
	}
	if len(components) > 0 {
		return &goast.BasicLit{
			Value: fmt.Sprintf("`%s`", strings.Join(components, ",")),
			Kind:  gotoken.STRING,
		}
	} else {
		return nil
	}
}

// TODO really lookup values from module and imports
func (ctx *moduleContext) lookupValue(val Value) Value {
	return val
}

func (ctx *moduleContext) lookupType(reference TypeReference) Type {
	if assignment := ctx.lookupContext.AssignmentList.GetType(reference.Name()); assignment != nil {
		return assignment.Type
	} else if usefulType, ok := USEFUL_TYPES[reference.Name()]; ok {
		return usefulType
	} else {
		ctx.appendError(errors.New(fmt.Sprintf("Can not resolve Type Reference %v", reference.Name())))
		return nil
	}
}
