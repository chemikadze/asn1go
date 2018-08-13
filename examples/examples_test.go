package examples

import (
	"github.com/chemikadze/asn1go"
	"io/ioutil"
	"testing"
)

func testExampleParsing(t *testing.T, filename string) *asn1go.ModuleDefinition {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("Failed to read file: %s", err.Error())
	}
	str := string(content)
	def, err := asn1go.ParseString(str)
	if err != nil {
		t.Fatalf("Failed to parse %v\n\nExpected nil error, got %v", str, err.Error())
	}
	return def
}

func TestParseKerberos(t *testing.T) {
	testExampleParsing(t, "rfc4120.asn1")
}

func TestParseSNMP(t *testing.T) {
	testExampleParsing(t, "rfc1157.asn1")
}

func TestParseSNMPSMI(t *testing.T) {
	testExampleParsing(t, "rfc1155.asn1")
}
