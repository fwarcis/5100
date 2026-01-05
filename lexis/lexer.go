package lexis

import (
	"5100/lexis/lexstates"
	"5100/lexis/lextypes"
)

type Lexer struct {
	state *lexstates.State
	ctx   lexstates.ParserContext
}

func NewLexer(text string, state lexstates.State) *Lexer {
	l := &Lexer{state: &state}
	l.ctx = lexstates.ParserContext{
		State:    &state,
		Runes:    []rune(text),
		Position: 0,
	}
	return l
}

func (l *Lexer) Parse() (tokens []lextypes.Token, err error) {
	for err == nil && l.ctx.HasNext() {
		toks, e := l.state.Parse(&l.ctx)
		err = e
		tokens = append(tokens, toks...)
	}
	return tokens, err
}
