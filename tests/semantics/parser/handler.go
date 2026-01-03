package parser_test

import (
	"testing"

	"5100/lexis/lextypes"
	"5100/semantics"
	"5100/tests"
)

type Expected struct {
	Expression semantics.Expression
	Error      error
}

type TestOfParser = tests.Test[[]lextypes.Token, Expected]

func HandleTestOfParser(t *testing.T, test TestOfParser, testN int) {
	expr := semantics.Parse(test.Input)
	if expr != test.Expected.Expression {
		test.WantGotError(testN, test.Expected, expr)
	}
}
