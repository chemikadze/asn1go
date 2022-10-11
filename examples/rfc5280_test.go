package examples

import (
	"encoding/asn1"
	"encoding/pem"
	"io/ioutil"
	"testing"
)

// Has to use big.Int integers due to SerialNumber size.

//go:generate go run ../cmd/asn1go/main.go -default-integer-repr big.Int -package examples rfc5280.asn1 rfc5280_generated.go

func TestX509Declarations(t *testing.T) {
	var _ Certificate
	var _ CertificateList
}

func TestX509Parsing(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/_.google.crt")
	if err != nil {
		t.Fatalf("Failed to read crt: %v", err)
	}
	block, _ := pem.Decode(data)
	var cert Certificate
	if _, err := asn1.Unmarshal(block.Bytes, &cert); err != nil {
		t.Fatalf("Failed to parse certificate: %v", err)
	}
	t.Logf("%+v", cert)
}
