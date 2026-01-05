package syntax

import (
	"5100/lexis/lextypes"
)

type Node interface {
	Token() lextypes.Token
}

type Value struct {
	value lextypes.Token
}

func (v *Value) Token() lextypes.Token {
	return v.value
}

type Unary struct {
	Value
	Next Node
}

type Binary struct {
	Value
	Left  Node
	Right Node
}
