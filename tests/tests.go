package tests

import (
	"fmt"
	"testing"
)

type Test[I any, E any] struct {
	Name     string
	Input    I
	Expected E
}

func (test *Test[I, E]) WantGotError(position int, want any, got any) string {
	return fmt.Sprintf(
		"%v: at %d:\nwant  %s\ngot   %s\n",
		test.Input, position, want, got,
	)
}

func (test *Test[I, E]) ModuleError(text string) string {
	return fmt.Sprintf("%v:\n%s\n", test.Input, text)
}

type Testing[Inp any, Exp any] struct {
	t     *testing.T
	tests []Test[Inp, Exp]
}

func NewTesting[Inp any, Exp any](t *testing.T, tests []Test[Inp, Exp]) Testing[Inp, Exp] {
	return Testing[Inp, Exp]{t, tests}
}

func (ttng *Testing[Inp, Exp]) Run(f func(test Test[Inp, Exp], position int)) {
	for i, test := range ttng.tests {
		name := test.Name
		if name == "" {
			name = fmt.Sprintf("%v", test.Input)
		}
		ttng.t.Run(name, func(t *testing.T) {
			f(test, i)
		})
	}
}
