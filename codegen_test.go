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

func TestDeclPrimitiveTypesSyntax(t *testing.T) {
	m := ModuleDefinition{
		ModuleIdentifier: ModuleIdentifier{Reference: "My-ASN1-ModuleName"},
		ModuleBody: ModuleBody{
			AssignmentList: AssignmentList{
				TypeAssignment{TypeReference("MyBool"), BooleanType{}},
				TypeAssignment{TypeReference("MyInt"), IntegerType{}},
				TypeAssignment{TypeReference("MyString"), CharacterStringType{}},
				TypeAssignment{TypeReference("MyOctetString"), OctetStringType{}},
				TypeAssignment{TypeReference("MyReal"), RealType{}},
			},
		},
	}
	expected := `package My_ASN1_ModuleName

type MyBool bool
type MyInt int64
type MyString string
type MyOctetString []byte
type MyReal float64
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

func TestDeclSequenceTypeSyntax(t *testing.T) {
	m := ModuleDefinition{
		ModuleIdentifier: ModuleIdentifier{Reference: "My-ASN1-ModuleName"},
		ModuleBody: ModuleBody{
			AssignmentList: AssignmentList{
				TypeAssignment{TypeReference("MySequence"), SequenceType{Components: ComponentTypeList{
					NamedComponentType{NamedType: NamedType{
						Identifier: Identifier("myIntField"),
						Type:       IntegerType{},
					}},
					NamedComponentType{NamedType: NamedType{
						Identifier: Identifier("myStructField"),
						Type: SequenceType{Components: ComponentTypeList{
							NamedComponentType{NamedType: NamedType{
								Identifier: Identifier("myOctetString"),
								Type:       OctetStringType{},
							}},
						}},
					}},
				}}},
			},
		},
	}
	expected := `package My_ASN1_ModuleName

type MySequence struct {
	myIntField	int64
	myStructField	struct {
		myOctetString []byte
	}
}
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

func TestDeclSequenceOFTypeSyntax(t *testing.T) {
	m := ModuleDefinition{
		ModuleIdentifier: ModuleIdentifier{Reference: "My-ASN1-ModuleName"},
		ModuleBody: ModuleBody{
			AssignmentList: AssignmentList{
				TypeAssignment{TypeReference("MySequenceOfInt"), SequenceOfType{IntegerType{}}},
				TypeAssignment{TypeReference("MySequenceOfSequence"), SequenceOfType{SequenceType{Components: ComponentTypeList{
					NamedComponentType{NamedType: NamedType{
						Identifier: Identifier("myIntField"),
						Type:       IntegerType{},
					}}},
				}}},
			},
		},
	}
	expected := `package My_ASN1_ModuleName

type MySequenceOfInt []int64
type MySequenceOfSequence []struct {
	myIntField int64
}
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
