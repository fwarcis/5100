package semantics

import (
	"strconv"

	"5100/lexis/lextypes"
	"5100/syntax"
)

func Evaluate(expr syntax.Node) (float64, error) {
	if expr == nil {
		return 0, &NilNodeError{}
	}

	tok := expr.Token()
	switch node := expr.(type) {
	case *syntax.Unary:
		if tok.Type == lextypes.NumberType {
			val, err := strconv.ParseFloat(tok.Value, 64)
			return val, err
		}
	case *syntax.Binary:
		operate := binOpValsFuncs[lextypes.BinOpValue(tok.Value)]
		if node.Left == nil || node.Right == nil {
			return 0, &NilNodeError{}
		}
		leftVal, err := Evaluate(node.Left)
		if err != nil {
			return 0, err
		}
		rightVal, err := Evaluate(node.Right)
		if err != nil {
			return 0, err
		}
		return operate(leftVal, rightVal), nil
	}
	return 0, &UnexpectedNodeTypeError{}
}
