package asn1go

import (
	"testing"
)

func testNotFails(t *testing.T, str string) {
	_, err := ParseString(str)
	if err != nil {
		t.Errorf("Expected nil error, got %v on %v", err.Error(), str)
	}
}

func TestParseMinimalModule(t *testing.T) {
	testNotFails(t, "MyModule DEFINITIONS ::= BEGIN END")
	testNotFails(t, "MyModule { mymodule } DEFINITIONS ::= BEGIN END")
	testNotFails(t, "MyModule DEFINITIONS IMPLICIT TAGS ::= BEGIN END")
	testNotFails(t, "MyModule DEFINITIONS EXTENSIBILITY IMPLIED ::= BEGIN END")
}

