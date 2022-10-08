package asn1go

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func testNotFails(t *testing.T, str string) *ModuleDefinition {
	def, err := ParseString(str)
	if err != nil {
		t.Fatalf("Failed to parse %v\n\nExpected nil error, got %v", str, err.Error())
	}
	return def
}

func TestParseMinimalModule(t *testing.T) {
	var r *ModuleDefinition
	testNotFails(t, "MyModule DEFINITIONS ::= BEGIN END")
	testNotFails(t, "MyModule { mymodule } DEFINITIONS ::= BEGIN END")
	r = testNotFails(t, "MyModule DEFINITIONS IMPLICIT TAGS ::= BEGIN END")
	if r.TagDefault != TAGS_IMPLICIT {
		t.Error("IMPLICIT TAGS should set the flag")
	}
	r = testNotFails(t, "MyModule DEFINITIONS EXTENSIBILITY IMPLIED ::= BEGIN END")
	if r.ExtensibilityImplied != true {
		t.Error("EXTENSIBILITY IMPLIED should set the flag")
	}
}

func TestParseImports(t *testing.T) {
	content := `
	RFC1157-SNMP DEFINITIONS ::= BEGIN
		IMPORTS
			ObjectName, ObjectSyntax, NetworkAddress, IpAddress, TimeTicks
		  		FROM RFC1155-SMI;

		MyString ::= CHARACTER STRING  -- AssignmentList can't be empty (?)
	END
	`
	expected := []SymbolsFromModule{
		{
			Module: GlobalModuleReference{Reference: "RFC1155-SMI"},
			SymbolList: []Symbol{
				TypeReference("ObjectName"), TypeReference("ObjectSyntax"), TypeReference("NetworkAddress"),
				TypeReference("IpAddress"), TypeReference("TimeTicks"),
			},
		},
	}
	r := testNotFails(t, content)
	if es, rs := fmt.Sprintf("%+v", expected), fmt.Sprintf("%+v", r.ModuleBody.Imports); es != rs {
		t.Errorf("Imports did not match:\n exp: %v\n got: %v", es, rs)
	}
}

func TestDefinitiveIdentifier(t *testing.T) {
	content := `
	KerberosV5Spec2 {
        iso(1) identified-organization(3) dod(6)
        nameform
        42 --numberform
        mixedform(88)
	} DEFINITIONS EXPLICIT TAGS ::= BEGIN
	END
	`
	r := testNotFails(t, content)
	if r.ModuleIdentifier.Reference != "KerberosV5Spec2" {
		t.Errorf("Expected reference KerberosV5Spec2 to be parsed, got '%v'", r.ModuleIdentifier.Reference)
	}
	if len(r.ModuleIdentifier.DefinitiveIdentifier) != 6 {
		t.Errorf("Expected 6 segments to be parsed, got %v", len(r.ModuleIdentifier.DefinitiveIdentifier))
	}
	expected := []DefinitiveObjIdComponent{
		{"iso", 1},
		{"identified-organization", 3},
		{"dod", 6},
		{Name: "nameform"},
		{"", 42},
		{"mixedform", 88},
	}
	for i, el := range r.ModuleIdentifier.DefinitiveIdentifier {
		expectedEl := expected[i]
		if el.Name != expectedEl.Name {
			t.Errorf("Expected %v component '%v' got '%v'", i, el.Name, expectedEl.Name)
		}
		if el.Id != expectedEl.Id {
			t.Errorf("Expected %v component '%v' got '%v'", i, el.Id, expectedEl.Id)
		}
	}
}

func TestValueAssignmentOID(t *testing.T) {
	content := `
	KerberosV5Spec2 DEFINITIONS ::= BEGIN
		id-krb5         OBJECT IDENTIFIER ::= {
    	    name-form
    	    42  --number-form
    	    name-and-number-form(77)
		}
	END
	`
	r := testNotFails(t, content)
	assignments := r.ModuleBody.AssignmentList
	if len(assignments) != 1 {
		t.Fatalf("Expected 1 assignment to be parsed, got %v", len(assignments))
	}
	krb := assignments.GetValue("id-krb5")
	if krb == nil {
		t.Fatal("Expected assignment with name id-krb5 to exist, got nil")
	}
	if krb.ValueReference.Name() != "id-krb5" {
		t.Errorf("Expected assignment LHS to be id-krb5, got %v", krb.ValueReference.Name())
	}
	if krb.Type != (ObjectIdentifierType{}) {
		t.Errorf("Expected value to be of OID type, got %v", krb.Type)
	}
	switch v := krb.Value.(type) {
	case ObjectIdentifierValue:
		if v.Type() != krb.Type {
			t.Errorf("Expected assignment value to have same type as assignment itself, got %v != %v", v.Type(), krb.Type)
		}
		expected := []ObjectIdElement{
			{Name: "name-form"},
			{Id: 42},
			{Name: "name-and-number-form", Id: 77},
		}
		if len(expected) != len(v) {
			t.Fatalf("Expected %v elements, got %v", len(expected), len(v))
		}
		for i, el := range v {
			expectedEl := expected[i]
			if el != expectedEl {
				t.Errorf("Expected %v element to be %v, got %v", i, expectedEl, el)
			}
		}
	default:
		t.Errorf("Expected ObjectIdentifierValue, got %t", v)
	}
	// TODO test DefinedValue
}

func testReal(t *testing.T, input Real, expectedValue Real) {
	if input != expectedValue {
		t.Errorf("Expected real value to be '%v' to be read, got '%v'", expectedValue, input)
	}
}

func TestRealBuilder(t *testing.T) {
	testReal(t, parseRealNumber(0, 0, 0), Real(0.0))
	testReal(t, parseRealNumber(1, 0, 0), Real(1.0))
	testReal(t, parseRealNumber(12345, 0, 0), Real(12345.0))
	testReal(t, parseRealNumber(12, 34, 0), Real(12.34))
	testReal(t, parseRealNumber(2, 346, 1), Real(23.46))
	testReal(t, parseRealNumber(23, 46, -1), Real(2.346))
}

func TestRangeTypeConstraint(t *testing.T) {
	content := `
	KerberosV5Spec2 DEFINITIONS ::= BEGIN
		Int32 ::= INTEGER (0..5 | 42^10..15)  -- note, UNION has lower precedence than INTERSECTION
	END
	`
	r := testNotFails(t, content)
	expectedType := ConstraintedType{
		Type: IntegerType{},
		Constraint: Constraint{
			ConstraintSpec: SubtypeConstraint{Unions{
				Intersections{
					{Elements: ValueRange{
						LowerEndpoint: RangeEndpoint{Value: Number(0)},
						UpperEndpoint: RangeEndpoint{Value: Number(5)}}}},
				Intersections{
					{Elements: SingleValue{Number(42)}},
					{Elements: ValueRange{
						LowerEndpoint: RangeEndpoint{Value: Number(10)},
						UpperEndpoint: RangeEndpoint{Value: Number(15)}}}}},
			},
		},
	}
	parsedAssignment := r.ModuleBody.AssignmentList.GetType("Int32")
	if parsedAssignment == nil {
		t.Fatal("Expected Int32 in assignments")
	}
	if reflect.TypeOf(parsedAssignment.Type) != reflect.TypeOf(expectedType) {
		t.Errorf("Expected %v got %v", expectedType, parsedAssignment)
	}
	parsedType := parsedAssignment.Type.(ConstraintedType)
	if reflect.TypeOf(parsedType.Type) != reflect.TypeOf(expectedType.Type) {
		t.Errorf("Expected type to be %v got %v", expectedType.Type, parsedType.Type)
	}
	parsedConstrant := parsedType.Constraint.ConstraintSpec.(SubtypeConstraint)
	expectedConstraint := expectedType.Constraint.ConstraintSpec.(SubtypeConstraint)
	if len(parsedConstrant) != len(expectedConstraint) {
		t.Errorf("Constraint length mismatch:\n exp %v\n got %v", expectedConstraint, parsedConstrant)
	}
	for i := range parsedConstrant {
		parsedUnions := parsedConstrant[i].(Unions)
		expectedUnions := expectedConstraint[i].(Unions)
		if len(parsedUnions) != len(expectedUnions) {
			t.Fatalf("Unions length mismatch:\n exp %v\n got %v", expectedUnions, parsedUnions)
		}
		for j := range parsedUnions {
			parsedInters := parsedUnions[j]
			expectedInters := expectedUnions[j]
			if len(parsedInters) != len(expectedInters) {
				t.Fatalf("Intersections length mismatch:\n exp %v\n got %v", expectedInters, parsedInters)
			}
			for k := range parsedInters {
				parsedIntElem := parsedInters[k]
				expectedIntElem := expectedInters[k]
				if parsedIntElem.Elements != expectedIntElem.Elements {
					t.Errorf("Intersection elements mismatch:\n exp %v\n got %v", expectedIntElem, parsedIntElem)
				}
			}
		}
	}
}

func firstConstraintElements(ct ConstraintedType) Elements {
	return ct.Constraint.ConstraintSpec.(SubtypeConstraint)[0].(Unions)[0][0].Elements
}

func asRestrictedString(elements Elements) RestrictedStringType {
	return elements.(TypeConstraint).Type.(RestrictedStringType)
}

func TestTypeTypeConstraint(t *testing.T) {
	content := `
	KerberosV5Spec2 DEFINITIONS ::= BEGIN
		KerberosString  ::= GeneralString (IA5String)
	END
	`
	expectedType := ConstraintedType{
		Type: RestrictedStringType{LexType: GeneralString},
		Constraint: Constraint{
			ConstraintSpec: SubtypeConstraint{Unions{Intersections{
				{Elements: TypeConstraint{RestrictedStringType{LexType: IA5String}}},
			}}},
		},
	}
	r := testNotFails(t, content)
	parsedAssignment := r.ModuleBody.AssignmentList.GetType("KerberosString")
	if parsedAssignment == nil {
		t.Fatal("Expected KerberosString in assignments")
	}
	parsedType := parsedAssignment.Type.(ConstraintedType)
	if parsedType.Type.(RestrictedStringType) != expectedType.Type.(RestrictedStringType) {
		t.Errorf("Expected %v got %v", expectedType.Type, parsedType.Type)
	}
	parsedElements := firstConstraintElements(parsedType)
	expectedElements := firstConstraintElements(expectedType)
	if asRestrictedString(parsedElements) != asRestrictedString(expectedElements) {
		t.Errorf("Expected %v got %v", expectedElements, parsedElements)
	}
}

func TestSequenceWithTagsAndSequenceOf(t *testing.T) {
	content := `
	KerberosV5Spec2 DEFINITIONS ::= BEGIN
		PrincipalName   ::= SEQUENCE {
				name-type       [0] Int32,
				name-string     [1] SEQUENCE OF KerberosString
		}
	END
	`
	expectedType := SequenceType{Components: ComponentTypeList{
		NamedComponentType{NamedType: NamedType{
			Identifier: Identifier("name-type"),
			Type:       TaggedType{Tag: Tag{ClassNumber: Number(0)}, Type: TypeReference("Int32")},
		}},
		NamedComponentType{NamedType: NamedType{
			Identifier: Identifier("name-string"),
			Type:       TaggedType{Tag: Tag{ClassNumber: Number(1)}, Type: SequenceOfType{TypeReference("KerberosString")}},
		}},
	}}
	r := testNotFails(t, content)
	parsedAssignment := r.ModuleBody.AssignmentList.GetType("PrincipalName")
	if parsedAssignment == nil {
		t.Fatal("Expected PrincipalName in assignments")
	}
	parsedType := parsedAssignment.Type.(SequenceType)
	if len(parsedType.Components) != len(expectedType.Components) {
		t.Fatalf("Expected %v components got %v", len(expectedType.Components), len(parsedType.Components))
	}
	for i := range parsedType.Components {
		expectedComponent := expectedType.Components[i].(NamedComponentType)
		parsedComponent := parsedType.Components[i].(NamedComponentType)
		if ei, pi := expectedComponent.NamedType.Identifier, parsedComponent.NamedType.Identifier; ei != pi {
			t.Errorf("Expected identifier %v got %v", ei, pi)
		}
	}
	// quick and dirty
	if es, ps := fmt.Sprintf("%+v", expectedType), fmt.Sprintf("%+v", parsedType); es != ps {
		t.Errorf("Repr mismatch:\n exp: %v\n got: %v", es, ps)
	}
}

func TestBitStringWithSizeConstraint(t *testing.T) {
	content := `
	KerberosV5Spec2 DEFINITIONS ::= BEGIN
		KerberosFlags   ::= BIT STRING (SIZE (32..MAX))
	END
	`
	expectedType := ConstraintedType{
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
	}
	r := testNotFails(t, content)
	parsedAssignment := r.ModuleBody.AssignmentList.GetType("KerberosFlags")
	if parsedAssignment == nil {
		t.Fatal("Expected KerberosFlags in assignments")
	}
	parsedType := parsedAssignment.Type
	// quick and dirty
	if es, ps := fmt.Sprintf("%+v", expectedType), fmt.Sprintf("%+v", parsedType); es != ps {
		t.Errorf("Repr mismatch:\n exp: %v\n got: %v", es, ps)
	}
}

func TestConstrainedSequence(t *testing.T) {
	content := `
	TestSpec DEFINITIONS ::= BEGIN
		CONSTRAINED-SEQUENCE ::= SEQUENCE SIZE (1..MAX) OF INTEGER
	END
	`
	expectedType := ConstraintedType{
		Type: SequenceOfType{IntegerType{}},
		Constraint: SingleElementConstraint(SizeConstraint{
			Constraint: SingleElementConstraint(ValueRange{
				LowerEndpoint: RangeEndpoint{Value: Number(1)},
				UpperEndpoint: RangeEndpoint{},
			}),
		}),
	}
	r := testNotFails(t, content)
	parsedAssignment := r.ModuleBody.AssignmentList.GetType("CONSTRAINED-SEQUENCE")
	if parsedAssignment == nil {
		t.Fatal("Expected CONSTRAINED-SEQUENCE in assignments")
	}
	parsedType := parsedAssignment.Type
	// quick and dirty
	if es, ps := fmt.Sprintf("%+v", expectedType), fmt.Sprintf("%+v", parsedType); es != ps {
		t.Errorf("Repr mismatch:\n exp: %v\n got: %v", es, ps)
	}
}

func TestChoiceType(t *testing.T) {
	content := `
	TestSpec DEFINITIONS ::= BEGIN
		PDUs ::=
			  CHOICE {
						  get-request
							  GetRequest-PDU,

						  get-next-request
							  GetNextRequest-PDU,

						  get-response
							  GetResponse-PDU,

						  set-request
							  SetRequest-PDU,

						  trap
							  Trap-PDU
					  }
	END
	`
	expectedType := ChoiceType{AlternativeTypeList: []NamedType{
		{Identifier("get-request"), TypeReference("GetRequest-PDU")},
		{Identifier("get-next-request"), TypeReference("GetNextRequest-PDU")},
		{Identifier("get-response"), TypeReference("GetResponse-PDU")},
		{Identifier("set-request"), TypeReference("SetRequest-PDU")},
		{Identifier("trap"), TypeReference("Trap-PDU")},
	}}
	r := testNotFails(t, content)
	parsedAssignment := r.ModuleBody.AssignmentList.GetType("PDUs")
	if parsedAssignment == nil {
		t.Fatal("Expected PDUs in assignments")
	}
	parsedType := parsedAssignment.Type
	// quick and dirty
	if es, ps := fmt.Sprintf("%+v", expectedType), fmt.Sprintf("%+v", parsedType); es != ps {
		t.Errorf("Repr mismatch:\n exp: %v\n got: %v", es, ps)
	}
}

func TestChoiceTypeExtension(t *testing.T) {
	content := `
	TestSpec DEFINITIONS ::= BEGIN
		PDUs ::=
			  CHOICE {
						  get-request
							  GetRequest-PDU,

						  get-next-request
							  GetNextRequest-PDU,

						  get-response
							  GetResponse-PDU,

						  set-request
							  SetRequest-PDU,

						  trap
							  Trap-PDU,
						  ...,
						  extra-choice
							  Extra-Type
					  }
	END
	`
	expectedType := ChoiceType{
		AlternativeTypeList: []NamedType{
			{Identifier("get-request"), TypeReference("GetRequest-PDU")},
			{Identifier("get-next-request"), TypeReference("GetNextRequest-PDU")},
			{Identifier("get-response"), TypeReference("GetResponse-PDU")},
			{Identifier("set-request"), TypeReference("SetRequest-PDU")},
			{Identifier("trap"), TypeReference("Trap-PDU")},
		},
		ExtensionTypes: []ChoiceExtension{
			NamedType{Identifier("extra-choice"), TypeReference("Extra-Type")},
		},
	}
	r := testNotFails(t, content)
	parsedAssignment := r.ModuleBody.AssignmentList.GetType("PDUs")
	if parsedAssignment == nil {
		t.Fatal("Expected PDUs in assignments")
	}
	parsedType := parsedAssignment.Type
	// quick and dirty
	if es, ps := fmt.Sprintf("%+v", expectedType), fmt.Sprintf("%+v", parsedType); es != ps {
		t.Errorf("Repr mismatch:\n exp: %v\n got: %v", es, ps)
	}
}

func TestRealValues(t *testing.T) {
	content := `
	TestSpec DEFINITIONS ::= BEGIN
		plusNum INTEGER ::= 123
		minusNum INTEGER ::= -123
		plusReal REAL ::= 123.4
		minusReal REAL ::= -1.234
		plusExp REAL ::= 1.234e3
		minusExp REAL ::= 1234e-3
	END
	`
	expectedDecls := AssignmentList{
		ValueAssignment{ValueReference("plusNum"), IntegerType{}, Number(123)},
		ValueAssignment{ValueReference("minusNum"), IntegerType{}, Number(-123)},
		ValueAssignment{ValueReference("plusReal"), RealType{}, Real(123.4)},
		ValueAssignment{ValueReference("minusReal"), RealType{}, Real(-1.234)},
		ValueAssignment{ValueReference("plusExp"), RealType{}, Real(1234.0)},
		ValueAssignment{ValueReference("minusExp"), RealType{}, Real(1.234)},
	}
	r := testNotFails(t, content)
	// quick and dirty
	if es, ps := fmt.Sprintf("%+v", r.ModuleBody.AssignmentList), fmt.Sprintf("%+v", expectedDecls); es != ps {
		t.Errorf("Repr mismatch:\n exp: %v\n got: %v", es, ps)
	}
}

func TestBooleanValues(t *testing.T) {
	content := `
	TestSpec DEFINITIONS ::= BEGIN
		true BOOLEAN ::= TRUE
		false BOOLEAN ::= FALSE
	END
	`
	expectedDecls := AssignmentList{
		ValueAssignment{ValueReference("true"), BooleanType{}, Boolean(true)},
		ValueAssignment{ValueReference("false"), BooleanType{}, Boolean(false)},
	}
	r := testNotFails(t, content)
	// quick and dirty
	if es, ps := fmt.Sprintf("%+v", r.ModuleBody.AssignmentList), fmt.Sprintf("%+v", expectedDecls); es != ps {
		t.Errorf("Repr mismatch:\n exp: %v\n got: %v", es, ps)
	}
}

func TestAnyType(t *testing.T) {
	content := `
	TestSpec DEFINITIONS ::= BEGIN
		AttributeValue ::= ANY
		AttributeValue2 ::= ANY DEFINED BY something
	END`
	expectedDecls := AssignmentList{
		TypeAssignment{TypeReference: "AttributeValue", Type: AnyType{}},
		TypeAssignment{TypeReference: "AttributeValue2", Type: AnyType{"something"}},
	}
	r := testNotFails(t, content)
	if diff := cmp.Diff(expectedDecls, r.ModuleBody.AssignmentList); diff != "" {
		t.Errorf("Module did not match expected, diff (-want, +got):\n%v", diff)
	}
}
