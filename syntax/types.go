package syntax

import (
	"log/slog"

	"5100/lexis/lextypes"
)

var BinOpFuncs = map[lextypes.BinOpValue]BinaryOperatorFunc{
	lextypes.PlusValue:  sum,
	lextypes.MinusValue: sub,
	lextypes.MulValue:   mul,
	lextypes.DivValue:   div,
}

func sum(a, b float64) float64 { return a + b }
func sub(a, b float64) float64 { return a - b }
func mul(a, b float64) float64 { return a * b }
func div(a, b float64) float64 { return a / b }

type Expression interface {
	Evaluate() float64
}

type Node interface {
	Get() lextypes.Token
}

type Value struct {
	value lextypes.Token
}

func (v *Value) Get() lextypes.Token {
	return v.value
}

type Unary struct {
	Value
	Next Node
}

type Binary struct {
	Value
	Left Node
	Right Node
}

type Number struct {
	value float64
}

func NewNumber(value float64) *Number {
	return &Number{value}
}

func (number *Number) Evaluate() float64 {
	slog.Debug("%d", slog.Float64("evaluate: ", number.value))
	return number.value
}

type BinaryOperatorFunc func(left, right float64) float64

type BinaryOperator struct {
	left, right Expression
	operate     BinaryOperatorFunc
}

func NewPlus(left, right Expression) *BinaryOperator {
	return &BinaryOperator{left, right, sum}
}

func NewMinus(left, right Expression) *BinaryOperator {
	return &BinaryOperator{left, right, sub}
}

func NewMultiplication(left, right Expression) *BinaryOperator {
	return &BinaryOperator{left, right, mul}
}

func NewDivision(left, right Expression) *BinaryOperator {
	return &BinaryOperator{left, right, div}
}

func (binOp *BinaryOperator) Evaluate() float64 {
	value := binOp.operate(binOp.left.Evaluate(), binOp.right.Evaluate())
	slog.Debug("%d", slog.Float64("evaluate: ", value))
	return value
}
