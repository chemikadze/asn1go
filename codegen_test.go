package asn1go

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func generateDeclarationsString(m ModuleDefinition) (string, error) {
	bufw := bytes.NewBufferString("")
	gen := NewCodeGenerator(GenParams{})
	err := gen.Generate(m, bufw)
	if err != nil {
		return "", err
	} else {
		return bufw.String(), nil
	}
}

func testModule(assignments AssignmentList) ModuleDefinition {
	return ModuleDefinition{
		ModuleIdentifier: ModuleIdentifier{Reference: "My-ASN1-ModuleName"},
		ModuleBody: ModuleBody{
			AssignmentList: assignments,
		},
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
	m := testModule(AssignmentList{
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
	})
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
	m := testModule(AssignmentList{
		TypeAssignment{TypeReference("MySequenceOfInt"), SequenceOfType{IntegerType{}}},
		TypeAssignment{TypeReference("MySequenceOfSequence"), SequenceOfType{SequenceType{Components: ComponentTypeList{
			NamedComponentType{NamedType: NamedType{
				Identifier: Identifier("myIntField"),
				Type:       IntegerType{},
			}}},
		}}},
	})
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

func TestTags(t *testing.T) {
	m := testModule(AssignmentList{
		TypeAssignment{TypeReference("MySequence"), SequenceType{Components: ComponentTypeList{
			NamedComponentType{NamedType: NamedType{
				Identifier: Identifier("myStringField"),
				Type:       RestrictedStringType{IA5String},
			}},
		}}},
	})
	expected := `package My_ASN1_ModuleName

type MySequence struct {
	MyStringField string ` + "`asn1:\"ia5\"`" + `
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

func TestTagType(t *testing.T) {
	testCases := []struct {
		name      string
		asnModule string
		goModule  string
	}{
		{
			name: "default (explicit)",
			asnModule: `
	TestSpec DEFINITIONS ::= BEGIN
		Struct ::= SEQUENCE {
			untagged BOOLEAN,
			default [1] BOOLEAN,
			explicit [2] EXPLICIT BOOLEAN,
			implicit [3] IMPLICIT BOOLEAN
		}
	END
	`,
			goModule: `package TestSpec

type Struct struct {
	Untagged	bool
	Default		bool	` + "`asn1:\"explicit,tag:1\"`" + `
	Explicit	bool	` + "`asn1:\"explicit,tag:2\"`" + `
	Implicit	bool	` + "`asn1:\"tag:3\"`" + `
}
`,
		},
		{
			name: "module explicit tags",
			asnModule: `
	TestSpec DEFINITIONS EXPLICIT TAGS ::= BEGIN
		Struct ::= SEQUENCE {
			untagged BOOLEAN,
			default [1] BOOLEAN,
			explicit [2] EXPLICIT BOOLEAN,
			implicit [3] IMPLICIT BOOLEAN
		}
	END
	`,
			goModule: `package TestSpec

type Struct struct {
	Untagged	bool
	Default		bool	` + "`asn1:\"explicit,tag:1\"`" + `
	Explicit	bool	` + "`asn1:\"explicit,tag:2\"`" + `
	Implicit	bool	` + "`asn1:\"tag:3\"`" + `
}
`,
		},
		{
			name: "module implicit tags",
			asnModule: `
	TestSpec DEFINITIONS IMPLICIT TAGS ::= BEGIN
		Struct ::= SEQUENCE {
			untagged BOOLEAN,
			default [1] BOOLEAN,
			explicit [2] EXPLICIT BOOLEAN,
			implicit [3] IMPLICIT BOOLEAN
		}
	END
	`,
			goModule: `package TestSpec

type Struct struct {
	Untagged	bool
	Default		bool	` + "`asn1:\"tag:1\"`" + `
	Explicit	bool	` + "`asn1:\"explicit,tag:2\"`" + `
	Implicit	bool	` + "`asn1:\"tag:3\"`" + `
}
`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := parseModule(t, tc.asnModule)
			expected := tc.goModule
			got, err := generateDeclarationsString(*m)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err.Error())
			}
			if diff := cmp.Diff(expected, got); diff != "" {
				t.Errorf("Generated module did not match expected, diff (-want, +got): %v", diff)
			}
		})
	}

}

func parseModule(t *testing.T, s string) *ModuleDefinition {
	def, err := ParseString(s)
	if err != nil {
		t.Fatalf("Failed to parse module: %v", err)
	}
	return def
}

func TestTime(t *testing.T) {
	m := testModule(AssignmentList{
		TypeAssignment{TypeReference("MyTimeType"), TypeReference("GeneralizedTime")},
		TypeAssignment{TypeReference("MySequence"), SequenceType{Components: ComponentTypeList{
			NamedComponentType{NamedType: NamedType{
				Identifier: Identifier("myTimeField"),
				Type:       TypeReference("MyTimeType"),
			}},
		}}},
	})
	expected := `package My_ASN1_ModuleName

import "time"

type MyTimeType time.Time
type MySequence struct {
	MyTimeField time.Time ` + "`asn1:\"generalized\"`" + `
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

func TestBitString(t *testing.T) {
	m := testModule(AssignmentList{
		TypeAssignment{TypeReference("MyBitStringType"), ConstraintedType{
			Type: BitStringType{},
			Constraint: Constraint{ConstraintSpec: SubtypeConstraint{
				Unions{Intersections{IntersectionElements{Elements: SizeConstraint{Constraint: Constraint{ConstraintSpec: SubtypeConstraint{
					Unions{Intersections{IntersectionElements{Elements: ValueRange{
						LowerEndpoint: RangeEndpoint{Value: Number(32)},
						UpperEndpoint: RangeEndpoint{},
					},
					}}},
				}},
				},
				}}},
			}},
		}},
		TypeAssignment{TypeReference("MyNestedBitStringType"), TypeReference("MyBitStringType")},
		TypeAssignment{TypeReference("MySequence"), SequenceType{Components: ComponentTypeList{
			NamedComponentType{NamedType: NamedType{
				Identifier: Identifier("myNestedBitStringField"),
				Type:       TypeReference("MyNestedBitStringType"),
			}},
			NamedComponentType{NamedType: NamedType{
				Identifier: Identifier("myBitStringField"),
				Type:       TypeReference("MyBitStringType"),
			}},
			NamedComponentType{NamedType: NamedType{
				Identifier: Identifier("bitStringField"),
				Type:       BitStringType{},
			}},
		}}},
	})
	expected := `package My_ASN1_ModuleName

import "encoding/asn1"

type MyBitStringType asn1.BitString
type MyNestedBitStringType asn1.BitString
type MySequence struct {
	MyNestedBitStringField	asn1.BitString
	MyBitStringField	asn1.BitString
	BitStringField		asn1.BitString
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
