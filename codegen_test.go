package asn1go

import (
	"bytes"
	"testing"
)

func TestDeclMinSynax(t *testing.T) {
	m := ModuleDefinition{
		ModuleIdentifier: ModuleIdentifier{Reference: "My-ASN1-ModuleName"},
	}
	expected := `package My_ASN1_ModuleName
`
	bufw := bytes.NewBufferString("")
	gen := NewCodeGenerator(GEN_DECLARATIONS)
	err := gen.Generate(m, bufw)
	if err != nil {
		t.Errorf("Unexpected error: %v", err.Error())
	}
	got := bufw.String()
	if got != expected {
		t.Errorf("Output did not match\n\nExp:\n`%v`\n\nGot:\n`%v`", expected, got)
	}
}
