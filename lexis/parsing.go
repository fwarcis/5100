package lexis

import (
	"5100/lexis/lexiserrs"
	"slices"
	"unicode"
)

var numberSigns = []string{"-", "+"}

type Lexer struct {
	runes []rune
	
	position int
	tokens []Token
	rn rune
	char string

	state parsingState
	numSign string
	digits string
}

func NewLexer(text string) *Lexer {
	lexer := &Lexer{
		runes: []rune(text),
		numSign: "+",
	}
	lexer.state = &stateSignOrDigit{lexer}
	
	return lexer
}

func (lex *Lexer) Parse() (*[]Token, error) {
	for lex.position = range lex.runes {
		lex.rn = lex.runes[lex.position]
		lex.char = string(lex.rn)

		if !lex.state.Valid() {
			return nil, &lexiserrs.ErrIllegalChar{Char: lex.char, Position: lex.position}
		} else if unexpected(lex.rn) {
			return nil, &lexiserrs.ErrUnexpectedChar{Char: lex.char, Position: lex.position}
		} 

		lex.state.Handle()
	}

	if lex.position+1 == len(lex.runes) {
		lex.tokens = append(
			lex.tokens,
			*NewNumber(string(lex.numSign) + lex.digits))
	}

	return &lex.tokens, nil
}

func unexpected(r rune) bool {
	return !unicode.IsDigit(r) &&
		!slices.Contains(binOpValues, BinOpValue(r))
}
