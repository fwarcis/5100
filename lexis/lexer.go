package lexis

type Lexer struct {
	tokens []Token
	state  *ParsingState
	ctx    parsingContext
}

func NewLexer(text string, state ParsingState) *Lexer {
	l := &Lexer{state: &state}
	l.ctx = parsingContext{
		State:    &state,
		Runes:    []rune(text),
		Position: 0,
	}
	return l
}

func (l *Lexer) Parse() (tokens []Token, err error) {
	for err == nil && l.ctx.HasNext() {
		toks, e := l.state.handle(&l.ctx)
		err = e
		tokens = append(tokens, toks...)
	}
	return tokens, err
}
