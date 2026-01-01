package lexis_test

import (
	"errors"
	"fmt"
	"testing"

	"5100/lexis"
	"5100/tests"
)

type ExpectedTokensAndError struct {
	Tokens []lexis.Token
	Error  error
}

func TestParse(t *testing.T) {
	ttng := tests.NewTesting(t, []tests.Test[string, *ExpectedTokensAndError]{
		{
			Input: "3+4",
			Expected: &ExpectedTokensAndError{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+3"),
					*lexis.NewPlus(),
					*lexis.NewNumber("+4"),
				},
				Error: nil,
			},
		},
		{
			Input: "-3+4",
			Expected: &ExpectedTokensAndError{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-3"),
					*lexis.NewPlus(),
					*lexis.NewNumber("+4"),
				},
				Error: nil,
			},
		},
		{
			Input: "+3-4",
			Expected: &ExpectedTokensAndError{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+3"),
					*lexis.NewMinus(),
					*lexis.NewNumber("+4"),
				},
				Error: nil,
			},
		},
		{
			Input: "-3-4",
			Expected: &ExpectedTokensAndError{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-3"),
					*lexis.NewMinus(),
					*lexis.NewNumber("+4"),
				},
				Error: nil,
			},
		},
		{
			Input: "3*4",
			Expected: &ExpectedTokensAndError{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+3"),
					*lexis.NewMultiplication(),
					*lexis.NewNumber("+4"),
				},
				Error: nil,
			},
		},
		{
			Input: "-3*4",
			Expected: &ExpectedTokensAndError{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-3"),
					*lexis.NewMultiplication(),
					*lexis.NewNumber("+4"),
				},
				Error: nil,
			},
		},
		{
			Input: "3*-4",
			Expected: &ExpectedTokensAndError{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+3"),
					*lexis.NewMultiplication(),
					*lexis.NewNumber("-4"),
				},
				Error: nil,
			},
		},
		{
			Input: "-3*-4",
			Expected: &ExpectedTokensAndError{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-3"),
					*lexis.NewMultiplication(),
					*lexis.NewNumber("-4"),
				},
				Error: nil,
			},
		},
		{
			Input: "3/4",
			Expected: &ExpectedTokensAndError{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+3"),
					*lexis.NewDivision(),
					*lexis.NewNumber("+4"),
				},
				Error: nil,
			},
		},
		{
			Input: "-3/4",
			Expected: &ExpectedTokensAndError{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-3"),
					*lexis.NewDivision(),
					*lexis.NewNumber("+4"),
				},
				Error: nil,
			},
		},
		{
			Input: "3/-4",
			Expected: &ExpectedTokensAndError{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+3"),
					*lexis.NewDivision(),
					*lexis.NewNumber("-4"),
				},
				Error: nil,
			},
		},
		{
			Input: "-3/-4",
			Expected: &ExpectedTokensAndError{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-3"),
					*lexis.NewDivision(),
					*lexis.NewNumber("-4"),
				},
				Error: nil,
			},
		},
	})

	ttng.Run(func(test tests.Test[string, *ExpectedTokensAndError], position int) {
		lexer := lexis.NewLexer(test.Input, *lexis.NewValueState())
		tokens, err := lexer.Parse()
		if !errors.Is(err, test.Expected.Error) {
			t.Error(test.UnexpectedErrError(err, test.Expected.Error),)
		}
		if len(tokens) != len(test.Expected.Tokens) {
			t.Error(test.WantGotError(
				position,
				fmt.Sprintf("length = %d", len(test.Expected.Tokens)),
				fmt.Sprintf("length = %d", len(tokens)),
			))
		}
		for i := range tokens {
			if tokens[i] != test.Expected.Tokens[i] {
				t.Error(test.WantGotError(
					position,
					test.Expected.Tokens[i],
					tokens[i]),
				)
			}
		}
	})
}

