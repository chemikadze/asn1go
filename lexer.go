package asn1go

import (
	"bufio"
	"errors"
	"fmt"
	"unicode"
	"io"
	"bytes"
	"strconv"
	"math"
)

type MyLexer struct {
	bufReader *bufio.Reader
	err       error
	result    AstNode

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
		if unicode.IsUpper(r) {
			lex.unreadRune()
			callback := func(s string) { lval.typeref = TypeReference(s) }
			return lex.consumeWord(callback, TYPEREFERENCE, "TYPE REFERENCE")
		} else if unicode.IsLower(r) {
			lex.unreadRune()
			callback := func(s string) { lval.identifier = Identifier(s) }
			return lex.consumeWord(callback, IDENTIFIER, "IDENTIFIER")
		} else if unicode.IsDigit(r) {
			lex.unreadRune()
			return lex.consumeNumberOrReal(lval, math.NaN())
		} else if r == '-' {
			return lex.consumeNumberOrReal(lval, -1)
		} else {
			fmt.Printf("!!! Skipped '%c'\n", r)

		}
	}
}

func (lex *MyLexer) consumeNumberOrReal(lval *yySymType, realStart float64) int {
	// worknig on this function at 11 PM was bad idea
	realValue := realStart
	fullRepr := ""
	if realStart == -1 {
		fullRepr += "-"
	}
	res := lex.consumeNumber(lval)
	if res != NUMBER {
		return res
	}
	realValue *= float64(int(lval.number))
	fullRepr = lval.numberRepr
	if lex.peekRune() == '.' {
		if math.IsNaN(realValue) {
			realValue = float64(int(lval.number))
		}
		lex.readRune()
		fullRepr += "."
		if unicode.IsDigit(lex.peekRune()) {
			res := lex.consumeNumber(lval)
			if res != NUMBER {
				return res
			}
			fullRepr += lval.numberRepr
			shift := float64(math.Pow10(int(math.Ceil(math.Log10(float64(lval.number))))))
			realValue = realValue + float64(lval.number) / shift
		}
	}
	if unicode.ToLower(lex.peekRune()) == 'e' {
		if math.IsNaN(realValue) {
			realValue = float64(int(lval.number))
		}
		r, _, _ := lex.readRune()
		fullRepr += string(r)
		exponent := 1
		possibleSignRune := lex.peekRune()
		if possibleSignRune == '-' {
			exponent = -1
			fullRepr += string(possibleSignRune)
			lex.readRune()
		}
		firstExponentRune := lex.peekRune()
		if unicode.IsDigit(firstExponentRune) {
			res := lex.consumeNumber(lval)
			if res != NUMBER {
				return res
			}
			exponent *= int(lval.number)
			fullRepr += lval.numberRepr
			realValue *= math.Pow10(exponent)
		} else {
			lex.Error(fmt.Sprintf("Expected exponent after '%v' got '%c'", fullRepr, firstExponentRune))
		}
	}
	if math.IsNaN(realValue) {
		return NUMBER
	} else {
		lval.real = Real(realValue)
		lval.numberRepr = fullRepr
		return REAL
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

func (lex *MyLexer) consumeWord(setter func(string), lexType int, lexName string) int {
	r, _, _ := lex.bufReader.ReadRune()
	acc := bytes.NewBufferString("")
	acc.WriteRune(r)
	lastR := r
	for {
		r, _, err := lex.bufReader.ReadRune()
		if err == io.EOF || isWhitespace(r)  {
			label := acc.String()
			setter(label)
			if label[len(label)-1] == '-' {
				lex.Error(fmt.Sprintf("%v can not end on hyphen, got %v", lexName, label))
				return -1
			}
			return lexType
		}
		if err != nil {
			lex.Error(fmt.Sprintf("Failed to read: %v", err.Error()))
			return -1
		}
		if !isIdentifierChar(r) {
			acc.WriteRune(r)
			lex.Error(fmt.Sprintf("Expected valid identifier char, got '%c' while reading '%v'", r, acc.String()))
			return -1
		}
		acc.WriteRune(r)
		if lastR == '-' && r == '-' {
			lex.Error(fmt.Sprintf("%v can not contain two hyphens in a row, got %v", lexName, acc.String()))
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
			lval.number = Number(i)
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
		return true;
	//VERTICAL TABULATION (11)
	case 11:
		return true;
	//FORM FEED (12)
	case 12:
		return true;
	//CARRIAGE RETURN (13)
	case 13:
		return true;
	default:
		return false;
	}
}

func isIdentifierChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-'
}
