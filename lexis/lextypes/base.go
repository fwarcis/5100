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
}

func (tok Token) String() string {
	return fmt.Sprintf(
		"Token{Type=%s, Value=%q}",
		tok.Type,
		tok.Value,
	)
}
