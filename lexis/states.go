package lexis

import (
	"slices"
	"unicode"
)

type parsingState interface {
	Handle()
	Valid() bool
}

type signOrDigitState struct {
	lex *Lexer
}

func (s *signOrDigitState) Handle() {
	if s.lex.numSign == "-" && s.lex.char == "-" ||
		s.lex.numSign == "+" && s.lex.char == "+" {
		s.lex.numSign = "+"
	} else if s.lex.numSign == "+" && s.lex.char == "-" ||
		s.lex.numSign == "-" && s.lex.char == "+" {
		s.lex.numSign = "-"
	} else if unicode.IsDigit(s.lex.rn) {
		s.lex.digits += s.lex.char
	}

	s.lex.state = &digitOrBinOpState{s.lex}
}

func (s *signOrDigitState) Valid() bool {
	return slices.Contains(numberSigns, string(s.lex.rn)) ||
			unicode.IsDigit(s.lex.rn)
}

type digitOrBinOpState struct {
	lex *Lexer
}

func (s *digitOrBinOpState) Handle() {
	if unicode.IsDigit(s.lex.rn) {
		s.lex.digits += s.lex.char
	} else if slices.Contains(binOpValues, BinOpValue(s.lex.rn)) {
		binOpVal := BinOpValue(s.lex.rn)
		if binOpVal == MinusValue {
			binOpVal = PlusValue
		}
		number := *NewNumber(string(s.lex.numSign) + s.lex.digits)
		s.lex.tokens = append(s.lex.tokens, number)
		s.lex.tokens = append(s.lex.tokens, *binOps[binOpVal]())

		s.lex.digits = ""
		if s.lex.char == "-" {
			s.lex.numSign = "-"
		} else {
			s.lex.numSign = "+"
		}
		s.lex.state = &signOrDigitState{s.lex}
	}
}

func (s *digitOrBinOpState) Valid() bool {
	return slices.Contains(binOpValues, BinOpValue(s.lex.rn)) ||
			unicode.IsDigit(s.lex.rn)
}
