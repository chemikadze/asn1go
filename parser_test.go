package asn1go

import (
	"testing"
)

func TestParseTypeReference(t *testing.T) {
	name :=  "MyTypeReference"
	result, err := ParseString(name)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err.Error())
	}
	expected := TypeReference(name)
	if result != expected {
		t.Errorf("Expected %v got %v", expected, result)
	}
}

func TestParseNumber(t *testing.T) {
	name :=  "12345"
	result, err := ParseString(name)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err.Error())
	}
	expected := Number(12345)
	if result != expected {
		t.Errorf("Expected %v got %v", expected, result)
	}
}
