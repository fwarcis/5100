package semantics

import "5100/lexis/lextypes"

type BinOpFunc func(left, right float64) float64

var BinOpValsFuncs = map[lextypes.BinOpValue]BinOpFunc{
	lextypes.PlusValue:  sum,
	lextypes.MinusValue: sub,
	lextypes.MulValue:   mul,
	lextypes.DivValue:   div,
}

func sum(a, b float64) float64        { return a + b }
func mul(a, b float64) float64        { return a * b }
func sub(left, right float64) float64 { return left - right }
func div(left, right float64) float64 { return left / right }
