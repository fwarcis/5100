package parser

import (
	"errors"
	"fmt"
	"testing"

	"5100/lexis"
	"5100/lexis/lexstates"
	"5100/lexis/lextypes"
	"5100/tests"
)

type Expected struct {
	Tokens []lextypes.Token
	Error  error
}
type TestOfParser = tests.Test[string, *Expected]

func HandleTestOfParser(t *testing.T, test TestOfParser, testN int) {
	l := lexis.NewLexer(test.Input, *lexstates.NewValueState())
	tokens, err := l.Parse()
	if !errors.Is(err, test.Expected.Error) {
		t.Error(test.UnexpectedErrError(testN, test.Expected.Error, err))
	}
	if len(tokens) != len(test.Expected.Tokens) {
		t.Error(test.WantGotError(
			testN,
			fmt.Sprintf("length = %d", len(test.Expected.Tokens)),
			fmt.Sprintf("length = %d", len(tokens)),
		))
		t.FailNow()
	}

	minToksToComp := min(len(tokens), len(test.Expected.Tokens))
	for i := range minToksToComp {
		if tokens[i] != test.Expected.Tokens[i] {
			t.Error(test.WantGotError(
				testN,
				test.Expected.Tokens[i],
				tokens[i]),
			)
		}
	}
}
