package asn1go

import (
	"testing"
	"bufio"
	"strings"
)

func lexForString(str string) *MyLexer {
	reader := bufio.NewReader(strings.NewReader(str))
	return &MyLexer{bufReader: reader}
}

func testLexem(t *testing.T, getter func(*yySymType) string, input string, expectedType int, output string) {
	str := input
	lex := lexForString(str)
	symType := &yySymType{}
	gotType := lex.Lex(symType)
	if lex.err != nil {
		t.Errorf("Expected nil error, got %v", lex.err)
	}
	if gotType != expectedType {
		t.Errorf("Expected %v token, got %v", expectedType, gotType)
	}
	expectedName := output
	if getter(symType) != expectedName {
		t.Errorf("Expected lexem '%v' to be read, got '%v'", expectedName, getter(symType))
	}
}

func testNumber(t *testing.T, input string, expectedValue Number) {
	str := input
	lex := lexForString(str)
	symType := &yySymType{}
	gotType := lex.Lex(symType)
	if lex.err != nil {
		t.Errorf("Expected nil error, got %v", lex.err)
	}
	expectedType := NUMBER
	if gotType != expectedType {
		t.Errorf("Expected %v token, got %v", expectedType, gotType)
	}
	if symType.number != expectedValue {
		t.Errorf("Expected lexem '%v' to be read, got '%v'", expectedValue, symType.number)
	}
}

func testReal(t *testing.T, input string, expectedValue Real) {
	str := input
	lex := lexForString(str)
	symType := &yySymType{}
	gotType := lex.Lex(symType)
	if lex.err != nil {
		t.Errorf("Expected nil error, got %v", lex.err)
	}
	expectedType := REAL
	if gotType != expectedType {
		t.Errorf("Expected %v token, got %v", expectedType, gotType)
	}
	if symType.real != expectedValue {
		t.Errorf("Expected lexem '%v' to be read, got '%v'", expectedValue, symType.real)
	}
}

func testError(t *testing.T, input string, expectedErr string) {
	str := input
	lex := lexForString(str)
	symType := &yySymType{}
	lex.Lex(symType)
	if lex.err.Error() != expectedErr {
		t.Errorf("Expected '%v' error, got '%v'", expectedErr, lex.err)
	}
}

func utr(t *yySymType) string {
	return string(t.typeref)
}

func ui(t *yySymType) string {
	return string(t.identifier)
}


func TestTypeReference(t *testing.T) {
	testLexem(t, utr, "MyTypeReference", TYPEREFERENCE, "MyTypeReference")
	testLexem(t, utr, "My-Type-Reference", TYPEREFERENCE, "My-Type-Reference")
	testError(t, "My--Type-Reference", "TYPE REFERENCE can not contain two hyphens in a row, got My--")
	testError(t, "MyTypeReference-", "TYPE REFERENCE can not end on hyphen, got MyTypeReference-")
	testError(t, "My$Type%Reference", "Expected valid identifier char, got '$' while reading 'My$'")
}

func TestIdentifier(t *testing.T) {
	testLexem(t, ui, "myIdentifier", IDENTIFIER, "myIdentifier")
	testLexem(t, ui, "my-Identifier", IDENTIFIER, "my-Identifier")
	testError(t, "my--Identifier", "IDENTIFIER can not contain two hyphens in a row, got my--")
	testError(t, "myIdentifier-", "IDENTIFIER can not end on hyphen, got myIdentifier-")
}

func TestSpacing(t *testing.T) {
	testLexem(t, ui, "   myIdentifier   ", IDENTIFIER, "myIdentifier")
}

func TestComments(t *testing.T) {
	testLexem(t, ui, "myIdentifier --thisisacomment", IDENTIFIER, "myIdentifier")

	testLexem(t, ui, `
	-- this is one comment
	-- this is second comment
	myIdentifier -- this is another comment
	-- this is trailing comment
	`, IDENTIFIER, "myIdentifier")

	testLexem(t, ui, "--comment1-- -- c-o-mm-ent2 -- myIdentifier --comment3-- ", IDENTIFIER, "myIdentifier")

	testLexem(t, ui, `
	/*
	this is a comment
	*/

	/* this is another comment */

	/*
	this
	/* is nested comment */
	*/

	myIdentifier
	`, IDENTIFIER, "myIdentifier")
}

func TestNumber(t *testing.T) {
	testNumber(t, "0", Number(0))
	testNumber(t, "1", Number(1))
	testNumber(t, "123", Number(123))
	testNumber(t, "12345", Number(12345))
}

func TestReal(t *testing.T) {
	testReal(t, "-1", Real(-1))
	testReal(t, "0.", Real(0.0))
	testReal(t, "1.", Real(1.0))
	testReal(t, "12345.", Real(12345.0))
	testReal(t, "12.34", Real(12.34))
	testReal(t, "2.346e1", Real(23.46))
	testReal(t, "23.46e-1", Real(2.346))
}