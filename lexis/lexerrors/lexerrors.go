package lexerrors

import (
	"fmt"
	"slices"
	"strings"

	"5100/lexis/lextypes"
)

func withPrefix(text string, r rune, position int) string {
	if r == -1 {
		return fmt.Sprintf("Lexic Error: End at %d: %s", position, text)
	}
	return fmt.Sprintf("Lexic Error: '%c' at %d: %s", r, position, text)
}

type UnexpectedTokenError struct {
	Position  int
	Rune rune
	Expecteds []lextypes.TokenType
}

func NewNumberExpectedError(position int, r rune) *UnexpectedTokenError {
	return &UnexpectedTokenError{
		Position:  position,
		Rune: r,
		Expecteds: []lextypes.TokenType{lextypes.NumberType},
	}
}

func (e *UnexpectedTokenError) Error() string {
	if len(e.Expecteds) == 0 {
		return withPrefix("Nothing to expect", e.Rune, e.Position)
	}
	expectedsBuilder := strings.Builder{}
	expectedsBuilder.WriteString(string(e.Expecteds[0]))
	for i, exp := range e.Expecteds[1:] {
		if i == len(e.Expecteds)-2 {
			expectedsBuilder.WriteString(" or " + string(exp))
			break
		}
		expectedsBuilder.WriteString(", " + string(exp))
	}
	return withPrefix(
		fmt.Sprintf("%s expected", expectedsBuilder.String()),
		e.Rune, e.Position)
}

func (e *UnexpectedTokenError) Is(target error) bool {
	switch targ := target.(type) {
	case *UnexpectedTokenError:
		return e.Position == targ.Position &&
			slices.Compare(e.Expecteds, targ.Expecteds) == 0
	}
	return false
}
