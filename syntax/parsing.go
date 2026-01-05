package syntax

import (
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
		return &Unary{tokValue: tokValue{currentTok}}
	}
	return &Binary{
		tokValue: tokValue{currentTok},
		Left:  Parse(tokens[:slicingPos]),
		Right: Parse(tokens[slicingPos+1:]),
	}
}

