package lexis

import (
	"log"
	"slices"
	"unicode"
)

type Lexer struct {
	tokens []Token
	state  *ParsingState
	ctx parsingContext
}

func NewLexer(text string, state ParsingState) *Lexer {
	l := &Lexer{state: &state}
	l.ctx = parsingContext{
		State: &state,
		Runes: []rune(text),
		Position: 0,
	}
	return l
}

func (l *Lexer) Parse() (tokens []Token, err error) {
	for err == nil && l.ctx.Position < len(l.ctx.Runes) {
		toks, e := l.state.handle(&l.ctx)
		err = e
		tokens = append(tokens, toks...)
		log.Println(l.ctx)
	}
	return tokens, err
}

func unexpected(r rune) bool {
	return !unicode.IsDigit(r) &&
		!slices.Contains(binOpValues, BinOpValue(r))
}
