package asn1go

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"unicode"
	"unicode/utf8"
)

var (
	RESERVED_WORDS map[string]int
)

type MyLexer struct {
	bufReader *bufio.Reader
	err       error
	result    *ModuleDefinition

	runeStack []rune
}

func (lex *MyLexer) Lex(lval *yySymType) int {
	for {
		r, _, err := lex.readRune()
		if err == io.EOF {
			return 0
		}
		if err != nil {
			lex.Error(fmt.Sprintf("Failed to read: %v", err.Error()))
			return -1
		}

		// fast forward cases
		if isWhitespace(r) {
			continue
		} else if r == '-' && lex.peekRune() == '-' {
			lex.skipLineComment()
			continue
		} else if r == '/' && lex.peekRune() == '*' {
			lex.skipBlockComment()
			continue
		}

		// parse lexem
		if unicode.IsLetter(r) {
			lex.unreadRune()
			content, err := lex.consumeWord()
			if err != nil {
				lex.Error(err.Error())
				return -1
			}
			if unicode.IsUpper(r) {
				code, exists := RESERVED_WORDS[content]
				if exists {
					return code
				} else {
					lval.name = content
					return TYPEORMODULEREFERENCE
				}
			} else {
				lval.name = content
				return VALUEIDENTIFIER
			}
		} else if unicode.IsDigit(r) {
			lex.unreadRune()
			return lex.consumeNumber(lval)
		} else if r == ':' && lex.peekRunes(2) == ":=" {
			lex.discard(2)
			return ASSIGNMENT
		} else if r == '.' && lex.peekRunes(2) == ".." {
			lex.discard(2)
			return ELLIPSIS
		} else if r == '.' && lex.peekRune() == '.' {
			lex.discard(1)
			return RANGE_SEPARATOR
		} else if r == '[' && lex.peekRune() == '[' {
			lex.discard(1)
			return LEFT_VERSION_BRACKETS
		} else if r == ']' && lex.peekRune() == ']' {
			lex.discard(1)
			return RIGHT_VERSION_BRACKETS
		} else {
			return lex.consumeSingleSymbol(r)
		}
	}
}

func (lex *MyLexer) consumeSingleSymbol(r rune) int {
	switch r {
	case '{':
		return OPEN_CURLY
	case '}':
		return CLOSE_CURLY
	case '<':
		return LESS
	case '>':
		return GREATER
	case ',':
		return COMMA
	case '.':
		return DOT
	case '(':
		return OPEN_ROUND
	case ')':
		return CLOSE_ROUND
	case '[':
		return OPEN_SQUARE
	case ']':
		return CLOSE_SQUARE
	case '-':
		return MINUS
	case ':':
		return COLON
	case '=':
		return EQUALS
	case '"':
		return QUOTATION_MARK
	case '\'':
		return APOSTROPHE
	case ' ': // TODO at which context it can be parsed?
		return SPACE
	case ';':
		return SEMICOLON
	case '@':
		return AT
	case '|':
		return PIPE
	case '!':
		return EXCLAMATION
	case '^':
		return CARET
	default:
		lex.Error(fmt.Sprintf("Unexpected character: %c", r))
		return -1
	}
}

func (lex *MyLexer) unreadRune() error {
	r := lex.bufReader.UnreadRune()
	if r != nil {
		panic(r.Error())
	}
	return r
}

func (lex *MyLexer) readRune() (rune, int, error) {
	r, n, err := lex.bufReader.ReadRune()
	return r, n, err
}

func (lex *MyLexer) peekRune() rune {
	r, _ := lex.peekRuneE()
	return r
}

func (lex *MyLexer) discard(n int) {
	lex.bufReader.Discard(n)
}

func (lex *MyLexer) peekRunes(n int) string {
	acc := bytes.NewBufferString("")
	pos := 0
	for n > 0 {
		for l := 1; l <= utf8.UTFMax; l++ {
			buf, err := lex.bufReader.Peek(pos + l)
			slice := buf[pos : pos+l]
			if pos+l <= len(buf) && utf8.FullRune(slice) {
				r, size := utf8.DecodeRune(slice)
				acc.WriteRune(r)
				pos += size
				n -= 1
				break
			}
			if err == io.EOF { // TODO if it is not a full rune, will swallow the error
				return acc.String()
			}
		}
	}
	return acc.String()
}

func (lex *MyLexer) peekRuneE() (rune, error) {
	r, _, err := lex.bufReader.ReadRune()
	if err == nil {
		lex.bufReader.UnreadRune()
	}
	return r, err
}

func (lex *MyLexer) skipLineComment() {
	lastIsHyphen := false
	for {
		r, _, err := lex.readRune()
		if isNewline(r) || err == io.EOF {
			return
		} else if r == '-' {
			if lastIsHyphen {
				return
			}
			lastIsHyphen = true
		} else {
			lastIsHyphen = false
		}
	}
}

func (lex *MyLexer) skipBlockComment() {
	lastIsOpeningSlash := false
	lastIsClosingStar := false
	for {
		r, _, err := lex.readRune()
		if err == io.EOF {
			return
		}
		if r == '/' {
			if lastIsClosingStar {
				return
			} else {
				lastIsOpeningSlash = true
				continue
			}
		} else if r == '*' {
			if lastIsOpeningSlash {
				lex.skipBlockComment()
			} else {
				lastIsClosingStar = true
				continue
			}
		}
		lastIsClosingStar = false
		lastIsOpeningSlash = false
	}
}

func (lex *MyLexer) consumeWord() (string, error) {
	r, _, _ := lex.bufReader.ReadRune()
	acc := bytes.NewBufferString("")
	acc.WriteRune(r)
	lastR := r
	for {
		r, _, err := lex.readRune()
		if err == io.EOF || isWhitespace(r) || !isIdentifierChar(r) {
			label := acc.String()
			if label[len(label)-1] == '-' {
				return "", errors.New(fmt.Sprintf("Token can not end on hyphen, got %v", label))
			}
			if err == nil {
				lex.unreadRune()
			}
			return label, nil
		}
		if err != nil {
			return "", errors.New(fmt.Sprintf("Failed to read: %v", err.Error()))
		}
		if !isIdentifierChar(r) {
			acc.WriteRune(r)
			return "", errors.New(fmt.Sprintf("Expected valid identifier char, got '%c' while reading '%v'", r, acc.String()))
		}
		acc.WriteRune(r)
		if lastR == '-' && r == '-' {
			return "", errors.New(fmt.Sprintf("Token can not contain two hyphens in a row, got %v", acc.String()))
		}
		lastR = r
	}
}

func (lex *MyLexer) consumeNumber(lval *yySymType) int {
	r, _, err := lex.bufReader.ReadRune()
	if err != nil {
		lex.Error(err.Error())
		return -1
	}
	acc := bytes.NewBufferString("")
	acc.WriteRune(r)
	for {
		r, _, err := lex.readRune()
		if err == io.EOF || !unicode.IsDigit(r) {
			if err == nil && !unicode.IsDigit(r) {
				lex.unreadRune()
			}
			repr := acc.String()
			i, err := strconv.Atoi(repr)
			if err != nil {
				lex.Error(fmt.Sprintf("Failed to parse number: %v", err.Error()))
				return -1
			}
			lval.numberRepr = repr
			lval.Number = Number(i)
			return NUMBER
		}
		if err != nil {
			lex.Error(fmt.Sprintf("Failed to read: %v", err.Error()))
			return -1
		}
		acc.WriteRune(r)
	}
}

func (lex *MyLexer) Error(e string) {
	lex.err = errors.New(e)
}

func isWhitespace(r rune) bool {
	switch x := int(r); x {
	//HORIZONTAL TABULATION (9)
	case 9:
		return true
	//LINE FEED (10)
	case 10:
		return true
	//VERTICAL TABULATION (11)
	case 11:
		return true
	//FORM FEED (12)
	case 12:
		return true
	//CARRIAGE RETURN (13)
	case 13:
		return true
	//SPACE (32)
	case 32:
		return true
	default:
		return false
	}
}

func isNewline(r rune) bool {
	switch x := int(r); x {
	//LINE FEED (10)
	case 10:
		return true
	//VERTICAL TABULATION (11)
	case 11:
		return true
	//FORM FEED (12)
	case 12:
		return true
	//CARRIAGE RETURN (13)
	case 13:
		return true
	default:
		return false
	}
}

func isIdentifierChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-'
}
