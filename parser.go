package asn1go

import (
	"bufio"
	"io"
	"math"
	"os"
	"strings"
)

func init() {
	yyErrorVerbose = true
}

// ParseString parses string containing ASN.1 definitions into ASN.1 AST.
func ParseString(str string) (*ModuleDefinition, error) {
	return ParseStream(strings.NewReader(str))
}

// ParseStream reads text of ASN.1 definitions from provided reader and parses it into ASN.1 AST.
func ParseStream(reader io.Reader) (*ModuleDefinition, error) {
	lex := &MyLexer{}
	lex.bufReader = bufio.NewReader(reader)
	yyParse(lex)
	if lex.err != nil {
		return nil, lex.err
	}
	return lex.result, nil
}

// ParseFile parses ASN.1 definition file into ASN.1 AST.
func ParseFile(name string) (*ModuleDefinition, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ParseStream(file)
}

func parseRealNumber(integer Number, fraction Number, exponent Number) Real {
	value := float64(integer)
	if fraction != 0 {
		shift := float64(math.Pow10(int(math.Ceil(math.Log10(float64(fraction))))))
		value += float64(fraction) / shift
	}
	if exponent != 0 {
		value *= math.Pow10(int(exponent))
	}
	return Real(value)
}
