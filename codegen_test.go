package asn1go

import (
	"bytes"
	"testing"
)

func generateDeclarationsString(m ModuleDefinition) (string, error) {
	bufw := bytes.NewBufferString("")
	gen := NewCodeGenerator(GEN_DECLARATIONS)
	err := gen.Generate(m, bufw)
	if err != nil {
		return "", err
	} else {
		return bufw.String(), nil
	}
}

func TestDeclMinSynax(t *testing.T) {
	m := ModuleDefinition{
		ModuleIdentifier: ModuleIdentifier{Reference: "My-ASN1-ModuleName"},
	}
	expected := `package My_ASN1_ModuleName
`
	got, err := generateDeclarationsString(m)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err.Error())
	}
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
	got, err := generateDeclarationsString(m)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err.Error())
	}
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
	MyIntField	int64
	MyStructField	struct {
		MyOctetString []byte
	}
}
`
	got, err := generateDeclarationsString(m)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err.Error())
	}
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
	MyIntField int64
}
`
	got, err := generateDeclarationsString(m)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err.Error())
	}
	if got != expected {
		t.Errorf("Output did not match\n\nExp:\n`%v`\n\nGot:\n`%v`", expected, got)
	}
}
