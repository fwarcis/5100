package lexstates

import (
	"slices"
	"unicode"

	"5100/lexis/lextypes"
)

var ValueStateExpecteds = []lextypes.TokenType{lextypes.NumberType}

func NewValueState() *State {
	return &State{
		Handlers: []Handler{
			&NumberHandler{},
		},
		Expecteds: ValueStateExpecteds,
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
	var lastRune rune = -1
	for _, r := range ctx.CurrentRunes() {
		if !unicode.IsDigit(r) {
			lastRune = r
			break
		}
		numVal += string(r)
		ctx.Position++
	}
	if numVal == "" || len(numVal) == 1 && hasPlusSign {
		return nil, false
	}

	*ctx.State = *NewSpaceState()
	if lastRune != ' ' {
		*ctx.State = *NewOperatorState()
	}
	return []lextypes.Token{*lextypes.NewNumber(numVal)}, true
}
