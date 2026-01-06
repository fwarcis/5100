package lexstates

import (
	"5100/lexis/lexerrors"
	"5100/lexis/lextypes"
)

type State struct {
	Handlers  []Handler
	Expecteds []lextypes.TokenType
}

func (s *State) Parse(ctx *ParserContext) ([]lextypes.Token, error) {
	tokens := []lextypes.Token{}
	nextTokPos := ctx.Position
	state := *s
	for i, h := range s.Handlers {
		toks, ok := h.Handle(ctx)
		tokens = toks
		if ok {
			ctx.PreviousState = &state
			return tokens, nil
		} else if i+1 == len(s.Handlers) {
			break
		}
		ctx.Position = nextTokPos
	}
	return tokens, &lexerrors.UnexpectedTokenError{
		Position:  ctx.Position,
		Rune:      ctx.Rune(),
		Expecteds: ctx.State.Expecteds,
	}
}

type Handler interface {
	Handle(ctx *ParserContext) (tokens []lextypes.Token, ok bool)
}

type ParserContext struct {
	State         *State
	PreviousState *State
	Runes         []rune
	Position      int
}

func (ctx *ParserContext) Rune() rune {
	if !ctx.HasRuneOnPosition() {
		return -1
	}
	return ctx.Runes[ctx.Position]
}

func (ctx *ParserContext) CurrentRunes() []rune {
	return ctx.Runes[ctx.Position:]
}

func (ctx *ParserContext) HasRuneOnPosition() bool {
	return ctx.Position < len(ctx.Runes)
}
