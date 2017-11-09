package asn1go

import (
	"testing"
	"io/ioutil"
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

func TestParseKerberos(t *testing.T) {
	content, err := ioutil.ReadFile("examples/rfc4120.asn1")
	if err != nil {
		t.Errorf("Failed to read file: %s", err.Error())
	}
	testNotFails(t, string(content))
}