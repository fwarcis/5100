package lexis

import "fmt"

type TokenType string

const (
	NumberType  TokenType = "Number"
	BinOpType             = "BinOpr"
)

var binOpsPriorities = map[string]int{
	"+": 1, "-": 1,
	"*": 2, "/": 2,
}

var binOpValues []string = func() []string {
	res := []string{}
	for op := range binOpsPriorities {
		res = append(res, op)
	}
	return res
}()

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

func NewNumber(value string) *Token {
	return &Token{
		Value:    value,
		Priority: 100,
		Type:     NumberType,
	}
}

func NewBinaryOperator(value string) *Token {
	return &Token{
		Value:    value,
		Type:     BinOpType,
		Priority: binOpsPriorities[value],
	}
}
