package semantics

import (
	"strconv"

	"5100/lexis"
)

func Parse(tokens []lexis.Token) Expression {
	relativePosition := 0
	firstMinPriorityToken := tokens[0]
	for pos, tok := range tokens {
		if tok.Priority < firstMinPriorityToken.Priority {
			firstMinPriorityToken = tok
			relativePosition = pos
		}
	}

	if firstMinPriorityToken.Type == lexis.NumberType {
		value, _ := strconv.ParseFloat(firstMinPriorityToken.Value, 64)
		return &Number{value: value}
	}
	return &BinaryOperator{
		left:    Parse(tokens[:relativePosition]),
		right:   Parse(tokens[relativePosition+1:]),
		operate: BinOpsFuncs[firstMinPriorityToken.Value],
	}
}
