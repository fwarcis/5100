package lextypes

import "fmt"

type TokenType string

const (
	NumberType TokenType = "Number"
	BinOpType  TokenType = "Binary Operator"
)

type BinOpValue string

const (
	PlusValue  BinOpValue = "+"
	MinusValue BinOpValue = "-"
	MulValue   BinOpValue = "*"
	DivValue   BinOpValue = "/"
)

var BinOps = map[BinOpValue](func() *Token){
	PlusValue:  NewPlus,
	MinusValue: NewMinus,
	MulValue:   NewMultiplication,
	DivValue:   NewDivision,
}

var BinOpValues []BinOpValue = *func() *[]BinOpValue {
	res := []BinOpValue{}
	for op := range BinOps {
		res = append(res, op)
	}
	return &res
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
		Value:    string(PlusValue),
		Type:     BinOpType,
		Priority: 1,
	}
}

func NewMinus() *Token {
	return &Token{
		Value:    string(MinusValue),
		Type:     BinOpType,
		Priority: 1,
	}
}

func NewMultiplication() *Token {
	return &Token{
		Value:    string(MulValue),
		Type:     BinOpType,
		Priority: 2,
	}
}

func NewDivision() *Token {
	return &Token{
		Value:    string(DivValue),
		Type:     BinOpType,
		Priority: 2,
	}
}
