package lexis

import "fmt"

type TokenType string

const (
	NumberType TokenType = "Number"
	BinOpType            = "BinOpr"
)

var binOps = map[string](func() *Token){
	"+": NewPlus,
	"-": NewMinus,
	"*": NewMultiplication,
	"/": NewDivision,
}

var binOpValues []string = func() []string {
	res := []string{}
	for op := range binOps {
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

func NewPlus() *Token {
	return &Token{
		Value:    "+",
		Type:     BinOpType,
		Priority: 1,
	}
}

func NewMinus() *Token {
	return &Token{
		Value:    "-",
		Type:     BinOpType,
		Priority: 1,
	}
}

func NewMultiplication() *Token {
	return &Token{
		Value:    "*",
		Type:     BinOpType,
		Priority: 2,
	}
}

func NewDivision() *Token {
	return &Token{
		Value:    "/",
		Type:     BinOpType,
		Priority: 2,
	}
}
