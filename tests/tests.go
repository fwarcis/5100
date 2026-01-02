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

func (t *Test[Inp, Exp]) WantGotError(testN int, want any, got any) string {
	return prefix(testN, wantGotRepr(want, got))
}

func (t *Test[Inp, Exp]) UnexpectedErrError(testN int, want, got error) string {
	return prefix(testN, wantGotRepr(
		fmt.Sprintf("error = {\n\tValue: %v\n\tType:  %T\n}", want, want),
		fmt.Sprintf("error = {\n\tValue: %v\n\tType:  %T\n}", got, got),
	))
}

func prefix(testN int, after string) string {
	return fmt.Sprintf("test #%d:\n%s", testN, after)
}

func wantGotRepr(want, got any) string {
	return fmt.Sprintf("want  %v\ngot   %v", want, got)
}

type Testing[Inp, Exp any] struct {
	T     *testing.T
	Tests []Test[Inp, Exp]
}

func NewTesting[Inp, Exp any](t *testing.T, tests []Test[Inp, Exp]) Testing[Inp, Exp] {
	return Testing[Inp, Exp]{t, tests}
}

func (ttng *Testing[Inp, Exp]) Run(handleTest func(t *testing.T, test Test[Inp, Exp], testN int)) {
	for i, test := range ttng.Tests {
		ttng.T.Run(test.NameOrInputRepr(), func(t *testing.T) {
			handleTest(t, test, i)
		})
	}
}
