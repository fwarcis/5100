package lexis

import (
	"slices"
	"unicode"
)

type parsingState interface {
	Handle()
	Valid() bool
}

type stateSignOrDigit struct {
	lex *Lexer
}

func (s *stateSignOrDigit) Handle() {
	if s.lex.numSign == '-' && s.lex.char == '-' ||
		s.lex.numSign == '+' && s.lex.char == '+' {
		s.lex.numSign = '+'
	} else if s.lex.numSign == '+' && s.lex.char == '-' ||
		s.lex.numSign == '-' && s.lex.char == '+' {
		s.lex.numSign = '-'
	} else if unicode.IsDigit(s.lex.char) {
		s.lex.digits += string(s.lex.char)
	}

	s.lex.state = &stateDigitOrBinOp{s.lex}
}

func (s *stateSignOrDigit) Valid() bool {
	return slices.Contains(numberSigns, string(s.lex.char)) ||
			unicode.IsDigit(s.lex.char)
}

type stateDigitOrBinOp struct {
	lex *Lexer
}

func (s *stateDigitOrBinOp) Handle() {
	if unicode.IsDigit(s.lex.char) {
		s.lex.digits += string(s.lex.char)
	} else if slices.Contains(binOpValues, BinOpValue(s.lex.char)) {
		binOpVal := BinOpValue(s.lex.char)
		if binOpVal == MinusValue {
			binOpVal = PlusValue
		}
		number := *NewNumber(string(s.lex.numSign) + s.lex.digits)
		s.lex.tokens = append(s.lex.tokens, number)
		s.lex.tokens = append(s.lex.tokens, *binOps[binOpVal]())

		s.lex.digits = ""
		if s.lex.char == '-' {
			s.lex.numSign = '-'
		} else {
			s.lex.numSign = '+'
		}
		s.lex.state = &stateSignOrDigit{s.lex}
	}
}

func (s *stateDigitOrBinOp) Valid() bool {
	return slices.Contains(binOpValues, BinOpValue(s.lex.char)) ||
			unicode.IsDigit(s.lex.char)
}
