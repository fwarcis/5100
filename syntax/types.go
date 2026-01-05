package syntax

import (
	"5100/lexis/lextypes"
)

type Node interface {
	Token() lextypes.Token
}

type tokValue struct {
	value lextypes.Token
}

func (v *tokValue) Token() lextypes.Token {
	return v.value
}

type Unary struct {
	tokValue
	Next Node
}

type Binary struct {
	tokValue
	Left  Node
	Right Node
}
