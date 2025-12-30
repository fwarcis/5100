package lexis

import (
	"5100/lexis/lexiserrs"
	"slices"
	"unicode"
)

var numberSigns = []string{"-", "+"}

type Lexer struct {
	runes []rune
	
	state parsingState
	char rune
	numSign rune
	digits string
	
	tokens []Token
}

func NewLexer(text string) *Lexer {
	l := &Lexer{
		runes: []rune(text),
		numSign: '+',
	}
	l.state = &stateSignOrDigit{l}
	return l
}

func (l *Lexer) Parse() (*[]Token, error) {
	pos := 0
	for i, r := range l.runes {
		l.char = r
		pos = i

		if !l.state.Valid() {
			return nil, &lexiserrs.ErrIllegalChar{Char: r, Position: i}
		} else if unexpected(r) {
			return nil, &lexiserrs.ErrUnexpectedChar{Char: r, Position: i}
		} 

		l.state.Handle()
	}

	if pos+1 == len(l.runes) {
		l.tokens = append(
			l.tokens,
			*NewNumber(string(l.numSign) + l.digits))
	}

	return &l.tokens, nil
}

func unexpected(r rune) bool {
	return !unicode.IsDigit(r) &&
		!slices.Contains(binOpValues, BinOpValue(r))
}
