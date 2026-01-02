package lexis

import (
	"fmt"
)

func sprefixf(format string, position int, a ...any) string {
	return fmt.Sprintf("lexis: at %d: "+format, position, a)
}

type UnexpectedTokenError struct {
	Position  int
	Expecteds []TokenType
}

func NewErrNumberExpected(position int) *UnexpectedTokenError {
	return &UnexpectedTokenError{
		Position:  position,
		Expecteds: []TokenType{NumberType},
	}
}

func (e *UnexpectedTokenError) Error() string {
	return sprefixf("%s expected", e.Position, e.Expecteds)
}
