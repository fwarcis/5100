package lexis

import (
	"slices"
	"unicode"
)

var numberSigns = []rune{'-', '+'}

type parsingContext struct {
	State    *ParsingState
	Runes    []rune
	Position int
}

func (ctx *parsingContext) Rune() rune {
	return ctx.Runes[ctx.Position]
}

func (ctx *parsingContext) CurrentRunes() []rune {
	return ctx.Runes[ctx.Position:]
}

func (ctx *parsingContext) HasNext() bool {
	return ctx.Position < len(ctx.Runes)
}

type ParsingState struct {
	handlers         []Handler
	expectedTokTypes []TokenType
}

func NewValueState() *ParsingState {
	return &ParsingState{
		handlers: []Handler{
			&NumberHandler{},
		},
		expectedTokTypes: []TokenType{NumberType},
	}
}

func NewOperatorState() *ParsingState {
	return &ParsingState{
		handlers: []Handler{
			&BinaryOperatorHandler{},
		},
		expectedTokTypes: []TokenType{BinOpType},
	}
}

func (s *ParsingState) handle(ctx *parsingContext) ([]Token, error) {
	tokens := []Token{}
	nextTokPos := ctx.Position
	for i, h := range s.handlers {
		toks, ok := h.handle(ctx)
		tokens = toks
		if ok {
			return tokens, nil
		} else if i+1 == len(s.handlers) {
			break
		}
		ctx.Position = nextTokPos
	}
	return tokens, &UnexpectedTokenError{
		Position:  ctx.Position,
		Expecteds: ctx.State.expectedTokTypes,
	}
}

type Handler interface {
	handle(ctx *parsingContext) (tokens []Token, ok bool)
}

type NumberHandler struct{}

func (*NumberHandler) handle(ctx *parsingContext) ([]Token, bool) {
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
	return []Token{*NewNumber(numVal)}, true
}

type BinaryOperatorHandler struct{}

func (*BinaryOperatorHandler) handle(ctx *parsingContext) ([]Token, bool) {
	for _, variant := range binOpValues {
		subtext := string(ctx.Runes[ctx.Position : ctx.Position+len(variant)])
		if subtext == string(variant) {
			*ctx.State = *NewValueState()
			ctx.Position += len(variant)
			tokens := []Token{*binOps[variant]()}
			if !ctx.HasNext() {
				return tokens, false
			}
			return tokens, true
		}
	}
	return nil, false
}
