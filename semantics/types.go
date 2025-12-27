package semantics

import (
	"log/slog"
)

type Expression interface {
	Evaluate() float64
}

type Number struct {
	value float64
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

func (binOp *BinaryOperator) Evaluate() float64 {
	value := binOp.operate(binOp.left.Evaluate(), binOp.right.Evaluate())
	slog.Debug("%d", slog.Float64("evaluate: ", value))
	return value
}

func Sum(a, b float64) float64 { return a + b }
func Sub(a, b float64) float64 { return a - b }
func Mul(a, b float64) float64 { return a * b }
func Div(a, b float64) float64 { return a / b }

var BinOpsFuncs = map[string]BinaryOperatorFunc{
	"+": Sum,
	"-": Sub,
	"*": Mul,
	"/": Div,
}
