package asn1go

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func testNotFails(t *testing.T, str string) *ModuleDefinition {
	def, err := ParseString(str)
	if err != nil {
		t.Errorf("Failed to parse %v\n\nExpected nil error, got %v", str, err.Error())
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

func TestParseKerberos(t *testing.T) {
	content, err := ioutil.ReadFile("examples/rfc4120.asn1")
	if err != nil {
		t.Errorf("Failed to read file: %s", err.Error())
	}
	testNotFails(t, string(content))
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

func TestTypeConstraint(t *testing.T) {
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
