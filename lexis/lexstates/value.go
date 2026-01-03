package lexstates

import (
	"slices"
	"unicode"

	"5100/lexis/lextypes"
)

func NewValueState() *State {
	return &State{
		Handlers: []Handler{
			&NumberHandler{},
		},
		Expecteds: []lextypes.TokenType{lextypes.NumberType},
	}
}

var numberSigns = []rune{'-', '+'}

type NumberHandler struct{}

func (*NumberHandler) Handle(ctx *ParserContext) ([]lextypes.Token, bool) {
	numVal := ""

	hasPlusSign := slices.Contains(numberSigns, ctx.Rune())

	if hasPlusSign {
		numVal += string(ctx.Rune())
		ctx.Position++
	}
	for _, r := range ctx.CurrentRunes() {
		if !unicode.IsDigit(r) {
			break
		}
		numVal += string(r)
		ctx.Position++
	}

	if numVal == "" || len(numVal) == 1 && hasPlusSign {
		return nil, false
	}
	*ctx.State = *NewOperatorState()
	return []lextypes.Token{*lextypes.NewNumber(numVal)}, true
}
