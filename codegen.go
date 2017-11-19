package asn1go

import (
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

type declCodeGen struct{}

func (declCodeGen) Generate(module ModuleDefinition, writer io.Writer) error {
	ast := &goast.File{
		Name: &goast.Ident{Name: goifyName(module.ModuleIdentifier.Reference)},
	}
	return goprint.Fprint(writer, gotoken.NewFileSet(), ast)
}

func goifyName(name string) string {
	return strings.Replace(name, "-", "_", -1)
}
