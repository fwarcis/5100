package lexstates

import (
	"5100/lexis/lextypes"
)

func NewOperatorState() *State {
	return &State{
		Handlers: []Handler{
			&BinaryOperatorHandler{},
		},
		Expecteds: []lextypes.TokenType{lextypes.BinOpType},
	}
}

type BinaryOperatorHandler struct{}

func (*BinaryOperatorHandler) Handle(ctx *ParserContext) ([]lextypes.Token, bool) {
	for _, variant := range lextypes.BinOpValues {
		subtext := string(ctx.Runes[ctx.Position : ctx.Position+len(variant)])
		if subtext == string(variant) {
			*ctx.State = *NewValueState()
			ctx.Position += len(variant)
			tokens := []lextypes.Token{*lextypes.BinOps[variant]()}
			if !ctx.HasNext() {
				return tokens, false
			}
			return tokens, true
		}
	}
	return nil, false
}
