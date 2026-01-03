package lexerrors

import (
	"fmt"
	"slices"

	"5100/lexis/lextypes"
)

func sprefixf(format string, position int, a ...any) string {
	return fmt.Sprintf("lexis: at %d: "+format, position, a)
}

type UnexpectedTokenError struct {
	Position  int
	Expecteds []lextypes.TokenType
}

func NewErrNumberExpected(position int) *UnexpectedTokenError {
	return &UnexpectedTokenError{
		Position:  position,
		Expecteds: []lextypes.TokenType{lextypes.NumberType},
	}
}

func (e *UnexpectedTokenError) Error() string {
	return sprefixf("%s expected", e.Position, e.Expecteds)
}

func (e *UnexpectedTokenError) Is(target error) bool {
	switch targ := target.(type) {
	case *UnexpectedTokenError:
		return e.Position == targ.Position &&
			slices.Compare(e.Expecteds, targ.Expecteds) == 0
	}
	return false
}
