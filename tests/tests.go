package tests

import (
	"fmt"
	"testing"
)

type Test[Inp, Exp any] struct {
	Name     string
	Input    Inp
	Expected Exp
}

func (t *Test[Inp, Exp]) NameOrInputRepr() string {
	if t.Name != "" {
		return t.Name
	}
	return fmt.Sprintf("%v", t.Input)
}

func wantGotRepr(want, got any) string {
	return fmt.Sprintf("want  %v\ngot  %v", want, got)
}

func (t *Test[Inp, Exp]) WantGotError(position int, want any, got any) string {
	return fmt.Sprintf(
		"%v: test #%d: \nwant  %s\ngot   %s\n",
		t.NameOrInputRepr(), position, want, got,
	)
}

func (t *Test[Inp, Exp]) UnexpectedErrError(want, got error) string {
	return fmt.Sprintf(
		"%v: unexpected error: \n%s",
		t.NameOrInputRepr(),
		wantGotRepr(
			fmt.Sprintf("error = '%v'", want),
			fmt.Sprintf("error = '%v'", got),
		),
	)
}

func (t *Test[Inp, Exp]) ModuleError(text string) string {
	return fmt.Sprintf("%v:\n%s\n", t.Input, text)
}

type Testing[Inp, Exp any] struct {
	t     *testing.T
	tests []Test[Inp, Exp]
}

func NewTesting[Inp, Exp any](t *testing.T, tests []Test[Inp, Exp]) Testing[Inp, Exp] {
	return Testing[Inp, Exp]{t, tests}
}

func (ttng *Testing[Inp, Exp]) Run(f func(test Test[Inp, Exp], position int)) {
	for i, test := range ttng.tests {
		ttng.t.Run(test.NameOrInputRepr(), func(t *testing.T) {
			f(test, i)
		})
	}
}
