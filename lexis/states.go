package lexis

import (
	"slices"
	"unicode"
)

var numberSigns = []rune{'-', '+'}

type parsingContext struct {
	State *ParsingState
	Runes []rune
	Position int
}

type ParsingState struct {
	handlers []Handler
	expectedTokTypes []TokenType
}

func NewValueState() *ParsingState {
	return &ParsingState{
		handlers: []Handler{
			&NumberHandler{},
		},
		expectedTokTypes: []TokenType{NumberType,},
	}
}

func NewOperatorState() *ParsingState {
	return &ParsingState{
		handlers: []Handler{
			&BinaryOperatorHandler{},
		},
		expectedTokTypes: []TokenType{BinOpType,},
	}
}

func (s *ParsingState) handle(ctx *parsingContext) ([]Token, error) {
	tokens := []Token{}
	for _, h := range s.handlers {
		tokens, ok := h.handle(ctx)
		if ok {
			return tokens, nil
		}
	}
	return tokens, &UnexpectedTokenError{
		Position: ctx.Position,
		Expecteds: ctx.State.expectedTokTypes,
	}
}


type Handler interface {
	handle(ctx *parsingContext) (tokens []Token, ok bool)
}

type NumberHandler struct {}

func (*NumberHandler) handle(ctx *parsingContext) ([]Token, bool) {
	numVal := ""
	alreadyWithSign := false
	hasFirstDigit := false
	for i, r := range ctx.Runes[ctx.Position:] {
		if i == 0 && !alreadyWithSign && slices.Contains(numberSigns, r) {
			alreadyWithSign = true
		} else if i == 0 && !unicode.IsDigit(r) && !slices.Contains(numberSigns, r) ||
		alreadyWithSign && !hasFirstDigit && slices.Contains(numberSigns, r) {
			ctx.Position += i
			return nil, false
		} else if i >= 1 && !unicode.IsDigit(r) {
			break
		}
		if unicode.IsDigit(r) && !hasFirstDigit {
			hasFirstDigit = true
		}
		numVal += string(r)
	}
	ctx.Position += len(numVal)
	if unicode.IsDigit(rune(numVal[0])) {
		numVal = "+" + numVal
	}
	*ctx.State = *NewOperatorState()
	return []Token{*NewNumber(numVal)}, true
}

type BinaryOperatorHandler struct {}

func (*BinaryOperatorHandler) handle(ctx *parsingContext) ([]Token, bool) {
	for _, variant := range binOpValues {
		subtext := string(ctx.Runes[ctx.Position:ctx.Position+len(variant)])
		if subtext == string(variant) {
			*ctx.State = *NewValueState()
			ctx.Position += len(variant)
			return []Token{*binOps[variant]()}, true
		}
	}
	return nil, false
}


