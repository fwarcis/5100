package syntax

import (
	"fmt"

	"5100/lexis/lextypes"
)

func priority(tok lextypes.Token) int {
	switch tok.Type {
	case lextypes.NumberType:
		return 100000
	case lextypes.BinOpType:
		return binOpValsPriorities[lextypes.BinOpValue(tok.Value)]
	}
	panic(fmt.Sprintf("unexpected token type '%s'", tok.Type))
}

var binOpValsPriorities = map[lextypes.BinOpValue]int{
	lextypes.PlusValue:  100,
	lextypes.MinusValue: 100,
	lextypes.MulValue:   200,
	lextypes.DivValue:   200,
}
