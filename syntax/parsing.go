package syntax

import (
	"fmt"

	"5100/lexis/lextypes"
)

func Parse(tokens []lextypes.Token) Node {
	slicingPos := 0
	currentTok := tokens[0]
	for pos, tok := range tokens {
		if priority(tok) < priority(currentTok) {
			currentTok = tok
			slicingPos = pos
		}
	}

	if currentTok.Type == lextypes.NumberType {
		return &Unary{Value: Value{currentTok}}
	}
	return &Binary{
		Value: Value{currentTok},
		Left:  Parse(tokens[:slicingPos]),
		Right: Parse(tokens[slicingPos+1:]),
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

