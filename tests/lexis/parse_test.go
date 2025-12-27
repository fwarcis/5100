package lexis_test

import (
	"fmt"
	"testing"

	"5100/lexis"
	"5100/tests"
)


func TestParse(t *testing.T) {
	ttng := tests.NewTesting(t, []tests.Test[string, *[]lexis.Token]{
		{
			Input: "3+4",
			Expected: &[]lexis.Token{
				*lexis.NewNumber("+3"),
				*lexis.NewPlus(),
				*lexis.NewNumber("+4"),
			},
		},
		{
			Input: "-3+4",
			Expected: &[]lexis.Token{
				*lexis.NewNumber("-3"),
				*lexis.NewPlus(),
				*lexis.NewNumber("+4"),
			},
		},
		{
			Input: "+3-4",
			Expected: &[]lexis.Token{
				*lexis.NewNumber("+3"),
				*lexis.NewPlus(),
				*lexis.NewNumber("-4"),
			},
		},
		{
			Input: "-3-4",
			Expected: &[]lexis.Token{
				*lexis.NewNumber("-3"),
				*lexis.NewPlus(),
				*lexis.NewNumber("-4"),
			},
		},
		{
			Input: "3*4",
			Expected: &[]lexis.Token{
				*lexis.NewNumber("+3"),
				*lexis.NewMultiplication(),
				*lexis.NewNumber("+4"),
			},
		},
		{
			Input: "-3*4",
			Expected: &[]lexis.Token{
				*lexis.NewNumber("-3"),
				*lexis.NewMultiplication(),
				*lexis.NewNumber("+4"),
			},
		},
		{
			Input: "3*-4",
			Expected: &[]lexis.Token{
				*lexis.NewNumber("+3"),
				*lexis.NewMultiplication(),
				*lexis.NewNumber("-4"),
			},
		},
		{
			Input: "-3*-4",
			Expected: &[]lexis.Token{
				*lexis.NewNumber("-3"),
				*lexis.NewMultiplication(),
				*lexis.NewNumber("-4"),
			},
		},
		{
			Input: "3/4",
			Expected: &[]lexis.Token{
				*lexis.NewNumber("+3"),
				*lexis.NewDivision(),
				*lexis.NewNumber("+4"),
			},
		},
		{
			Input: "-3/4",
			Expected: &[]lexis.Token{
				*lexis.NewNumber("-3"),
				*lexis.NewDivision(),
				*lexis.NewNumber("+4"),
			},
		},
		{
			Input: "3/-4",
			Expected: &[]lexis.Token{
				*lexis.NewNumber("+3"),
				*lexis.NewDivision(),
				*lexis.NewNumber("-4"),
			},
		},
		{
			Input: "-3/4",
			Expected: &[]lexis.Token{
				*lexis.NewNumber("-3"),
				*lexis.NewDivision(),
				*lexis.NewNumber("+4"),
			},
		},
	})

	ttng.Run(func(test tests.Test[string, *[]lexis.Token], position int) {
		lexer := lexis.NewLexer(test.Input)
		tokens, err := lexer.Parse()

		if err != nil {
			t.Error(test.ModuleError(err.Error()))
			return
		}
		if len(*tokens) != len(*test.Expected) {
			t.Error(test.WantGotError(
				position,
				fmt.Sprintf("len = %d", len(*test.Expected)),
				fmt.Sprintf("len = %d", len(*tokens)),
			))
			fmt.Println(tokens)
		}
		for i := range *tokens {
			if (*tokens)[i] != (*test.Expected)[i] {
				t.Error(test.WantGotError(
					position,
					(*test.Expected)[i],
					(*tokens)[i]),
				)
			}
		}
	})
}
