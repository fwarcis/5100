package semantics

import (
	"fmt"
	"strconv"

	"5100/lexis/lextypes"
)

func Parse(tokens []lextypes.Token) Expression {
	relativePosition := 0
	minPriorityTok := tokens[0]
	for pos, tok := range tokens {
		if priority(tok) < priority(minPriorityTok) {
			minPriorityTok = tok
			relativePosition = pos
		}
	}

	if minPriorityTok.Type == lextypes.NumberType {
		value, _ := strconv.ParseFloat(minPriorityTok.Value, 64)
		return &Number{value: value}
	}
	return &BinaryOperator{
		left:    Parse(tokens[:relativePosition]),
		right:   Parse(tokens[relativePosition+1:]),
		operate: BinOpFuncs[lextypes.BinOpValue(minPriorityTok.Value)],
	}
}

func priority(tok lextypes.Token) int {
	switch tok.Type {
	case lextypes.NumberType:
		return 100000
	case lextypes.BinOpType:
		return BinOpPriorities[lextypes.BinOpValue(tok.Value)]
	}
	panic(fmt.Sprintf("unexpected token type '%s'", tok.Type))
}

var BinOpPriorities = map[lextypes.BinOpValue]int{
	lextypes.PlusValue: 100,
	lextypes.MinusValue: 100,
	lextypes.MulValue: 200,
	lextypes.DivValue: 200,
}

