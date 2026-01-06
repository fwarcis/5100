package lexstates

import (
	"slices"

	"5100/lexis/lextypes"
)

func NewSpaceState() *State {
	return &State{
		Handlers:  []Handler{&WhitespaceHandler{}},
		Expecteds: []lextypes.TokenType{lextypes.NumberType, lextypes.BinOpType},
	}
}

type WhitespaceHandler struct{}

func (*WhitespaceHandler) Handle(ctx *ParserContext) ([]lextypes.Token, bool) {
	for _, r := range ctx.CurrentRunes() {
		if r != ' ' {
			break
		}
		ctx.Position++
	}
	if slices.Compare(ctx.PreviousState.Expecteds, OperatorStateExpecteds) == 0 {
		*ctx.State = *NewValueState()
	} else {
		*ctx.State = *NewOperatorState()
	}
	return nil, true
}
