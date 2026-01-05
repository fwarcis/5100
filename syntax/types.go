package syntax

import (
	"5100/lexis/lextypes"
)

type Node interface {
	Get() lextypes.Token
}

type Value struct {
	value lextypes.Token
}

func (v *Value) Get() lextypes.Token {
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
