package asn1go

import (
	"strings"
	"bufio"
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
