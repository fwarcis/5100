package semantics_test

import (
	"5100/lexis"
	"5100/semantics"
	"5100/tests"
	"testing"
)

func TestParse(t *testing.T) {
	ttng := tests.NewTesting(t, []tests.Test[*[]lexis.Token, semantics.Expression]{
		{
			Input: &[]lexis.Token{
				*lexis.NewNumber("+3"),
				*lexis.NewPlus(),
				*lexis.NewNumber("+4"),
			},
			Expected: semantics.NewPlus(
				semantics.NewNumber(+3),
				semantics.NewNumber(+4),
			),
		},
		{
			Input: &[]lexis.Token{
				*lexis.NewNumber("-3"),
				*lexis.NewPlus(),
				*lexis.NewNumber("+4"),
			},
			Expected: semantics.NewPlus(
				semantics.NewNumber(-3),
				semantics.NewNumber(+4),
			),
		},
		{
			Input: &[]lexis.Token{
				*lexis.NewNumber("+3"),
				*lexis.NewMinus(),
				*lexis.NewNumber("+4"),
			},
			Expected: semantics.NewMinus(
				semantics.NewNumber(+3),
				semantics.NewNumber(+4),
			),
		},
		{
			Input: &[]lexis.Token{
				*lexis.NewNumber("-3"),
				*lexis.NewMinus(),
				*lexis.NewNumber("+4"),
			},
			Expected: semantics.NewMinus(
				semantics.NewNumber(-3),
				semantics.NewNumber(+4),
			),
		},
		{
			Input: &[]lexis.Token{
				*lexis.NewNumber("+3"),
				*lexis.NewMinus(),
				*lexis.NewNumber("-4"),
			},
			Expected: semantics.NewMinus(
				semantics.NewNumber(+3),
				semantics.NewNumber(-4),
			),
		},
		{
			Input: &[]lexis.Token{
				*lexis.NewNumber("-3"),
				*lexis.NewMinus(),
				*lexis.NewNumber("-4"),
			},
			Expected: semantics.NewMinus(
				semantics.NewNumber(-3),
				semantics.NewNumber(-4),
			),
		},
		{
			Input: &[]lexis.Token{
				*lexis.NewNumber("+3"),
				*lexis.NewMultiplication(),
				*lexis.NewNumber("+4"),
			},
			Expected: semantics.NewMultiplication(
				semantics.NewNumber(+3),
				semantics.NewNumber(+4),
			),
		},
		{
			Input: &[]lexis.Token{
				*lexis.NewNumber("-3"),
				*lexis.NewMultiplication(),
				*lexis.NewNumber("+4"),
			},
			Expected: semantics.NewMultiplication(
				semantics.NewNumber(-3),
				semantics.NewNumber(+4),
			),
		},
		{
			Input: &[]lexis.Token{
				*lexis.NewNumber("+3"),
				*lexis.NewMultiplication(),
				*lexis.NewNumber("-4"),
			},
			Expected: semantics.NewMultiplication(
				semantics.NewNumber(+3),
				semantics.NewNumber(-4),
			),
		},
		{
			Input: &[]lexis.Token{
				*lexis.NewNumber("-3"),
				*lexis.NewMultiplication(),
				*lexis.NewNumber("-4"),
			},
			Expected: semantics.NewMultiplication(
				semantics.NewNumber(-3),
				semantics.NewNumber(-4),
			),
		},
		{
			Input: &[]lexis.Token{
				*lexis.NewNumber("+3"),
				*lexis.NewDivision(),
				*lexis.NewNumber("+4"),
			},
			Expected: semantics.NewDivision(
				semantics.NewNumber(+3),
				semantics.NewNumber(+4),
			),
		},
		{
			Input: &[]lexis.Token{
				*lexis.NewNumber("-3"),
				*lexis.NewDivision(),
				*lexis.NewNumber("+4"),
			},
			Expected: semantics.NewDivision(
				semantics.NewNumber(-3),
				semantics.NewNumber(+4),
			),
		},
		{
			Input: &[]lexis.Token{
				*lexis.NewNumber("+3"),
				*lexis.NewDivision(),
				*lexis.NewNumber("-4"),
			},
			Expected: semantics.NewDivision(
				semantics.NewNumber(+3),
				semantics.NewNumber(-4),
			),
		},
		{
			Input: &[]lexis.Token{
				*lexis.NewNumber("-3"),
				*lexis.NewDivision(),
				*lexis.NewNumber("-4"),
			},
			Expected: semantics.NewDivision(
				semantics.NewNumber(-3),
				semantics.NewNumber(-4),
			),
		},
	})

	ttng.Run(func(test tests.Test[*[]lexis.Token, semantics.Expression], position int) {
		expr := semantics.Parse(*test.Input)
		if expr != test.Expected {
			test.WantGotError(position, test.Expected, expr)
		}
	})
}
