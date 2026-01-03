package lextypes

import "fmt"

type TokenType string

const (
	NumberType TokenType = "Number"
	BinOpType  TokenType = "Binary Operator"
)

type Token struct {
	Value    string
	Type     TokenType
	Priority int
}

func (tok Token) String() string {
	return fmt.Sprintf(
		"Token{Type=%s, Value=%q, Priority=%d}",
		tok.Type,
		tok.Value,
		tok.Priority,
	)
}
