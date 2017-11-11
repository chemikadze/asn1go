package asn1go

import (
	"bufio"
	"math"
	"strings"
)

func ParseString(str string) (*ModuleDefinition, error) {
	lex := &MyLexer{}
	reader := strings.NewReader(str)
	lex.bufReader = bufio.NewReader(reader)
	yyParse(lex)
	if lex.err != nil {
		return nil, lex.err
	}
	return lex.result, nil
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
