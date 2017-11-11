package asn1go

import (
	"io/ioutil"
	"testing"
)

func testNotFails(t *testing.T, str string) *ModuleDefinition {
	def, err := ParseString(str)
	if err != nil {
		t.Errorf("Expected nil error, got %v on %v", err.Error(), str)
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

func TestParseKerberos(t *testing.T) {
	content, err := ioutil.ReadFile("examples/rfc4120.asn1")
	if err != nil {
		t.Errorf("Failed to read file: %s", err.Error())
	}
	testNotFails(t, string(content))
}
