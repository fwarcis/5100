package lexstates

import (
	"5100/lexis/lextypes"
)

var OperatorStateExpecteds = []lextypes.TokenType{lextypes.BinOpType}

func NewOperatorState() *State {
	return &State{
		Handlers: []Handler{
			&BinaryOperatorHandler{},
		},
		Expecteds: OperatorStateExpecteds,
	}
}

type BinaryOperatorHandler struct{}

func (*BinaryOperatorHandler) Handle(ctx *ParserContext) ([]lextypes.Token, bool) {
	for _, variant := range lextypes.BinOpValues {
		subtext := string(ctx.Runes[ctx.Position : ctx.Position+len(variant)])
		if subtext == string(variant) {
			ctx.Position += len(variant)
			*ctx.State = *NewValueState()
			if ctx.Rune() == ' ' {
				*ctx.State = *NewSpaceState()
			}
			tokens := []lextypes.Token{
				*lextypes.BinOpValsConstructors[variant](),
			}
			if !ctx.HasRuneOnPosition() {
				return tokens, false
			}
			return tokens, true
		}
	}
	return nil, false
}
