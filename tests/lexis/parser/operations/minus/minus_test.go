package plus_test

import (
	"testing"

	"5100/lexis/lexerrors"
	"5100/lexis/lextypes"
	"5100/tests"
	"5100/tests/lexis/parser"
)

func Test__NoSign(t *testing.T) {
	ttng := tests.NewTesting(t, []parser.TestOfParser{
		{
			Input: "0-0",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("0"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("0"),
				},
				Error: nil,
			},
		},
		{
			Input: "0-1",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("0"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("1"),
				},
				Error: nil,
			},
		},
		{
			Input: "1-0",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("1"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("0"),
				},
				Error: nil,
			},
		},
		{
			Input: "01-1",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("01"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("1"),
				},
				Error: nil,
			},
		},
		{
			Input: "1-01",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("1"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("01"),
				},
				Error: nil,
			},
		},
		{
			Input: "000-000",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("000"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("000"),
				},
				Error: nil,
			},
		},
		{
			Input: "0001-002",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("0001"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("002"),
				},
				Error: nil,
			},
		},
	})

	ttng.Run(parser.HandleTestOfParser)
}

func Test__Single_NoSign_Erroring(t *testing.T) {
	ttng := tests.NewTesting(t, []parser.TestOfParser{
		{
			Input: "0-",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("0"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(2),
			},
		},
		{
			Input: "1-",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("1"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(2),
			},
		},
		{
			Input: "01-",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("01"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(3),
			},
		},
		{
			Input: "000-",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("000"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(4),
			},
		},
		{
			Input: "0001-",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("0001"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(5),
			},
		},
		{
			Input: "000123-",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("000123"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(7),
			},
		},
	})

	ttng.Run(parser.HandleTestOfParser)
}

func Test__Plus_Signeds(t *testing.T) {
	ttng := tests.NewTesting(t, []parser.TestOfParser{
		{
			Input: "+0-0",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("+0"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("0"),
				},
				Error: nil,
			},
		},
		{
			Input: "+1-0",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("+1"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("0"),
				},
				Error: nil,
			},
		},
		{
			Input: "+0-1",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("+0"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("1"),
				},
				Error: nil,
			},
		},
		{
			Input: "+01-1",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("+01"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("1"),
				},
				Error: nil,
			},
		},
		{
			Input: "+1-01",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("+1"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("01"),
				},
				Error: nil,
			},
		},
		{
			Input: "+000-000",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("+000"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("000"),
				},
				Error: nil,
			},
		},
		{
			Input: "+9-0000",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("+9"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("0000"),
				},
				Error: nil,
			},
		},
		{
			Input: "+000123-4",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("+000123"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("4"),
				},
				Error: nil,
			},
		},
	})

	ttng.Run(parser.HandleTestOfParser)
}

func Test__Single_Plus_Signeds_Erroring(t *testing.T) {
	ttng := tests.NewTesting(t, []parser.TestOfParser{
		{
			Input: "+0-",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("+0"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(3),
			},
		},
		{
			Input: "+1-",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("+1"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(3),
			},
		},
		{
			Input: "+01-",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("+01"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(4),
			},
		},
		{
			Input: "+000-",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("+000"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(5),
			},
		},
		{
			Input: "+0001-",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("+0001"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(6),
			},
		},
		{
			Input: "+000123-",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("+000123"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(8),
			},
		},
	})

	ttng.Run(parser.HandleTestOfParser)
}

func Test__Minus_Signeds(t *testing.T) {
	ttng := tests.NewTesting(t, []parser.TestOfParser{
		{
			Input: "-0--0",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("-0"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("-0"),
				},
				Error: nil,
			},
		},
		{
			Input: "-1--0",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("-1"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("-0"),
				},
				Error: nil,
			},
		},
		{
			Input: "-0--1",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("-0"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("-1"),
				},
				Error: nil,
			},
		},
		{
			Input: "-01--1",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("-01"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("-1"),
				},
				Error: nil,
			},
		},
		{
			Input: "-1--01",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("-1"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("-01"),
				},
				Error: nil,
			},
		},
		{
			Input: "-000--000",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("-000"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("-000"),
				},
				Error: nil,
			},
		},
		{
			Input: "-9--0000",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("-9"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("-0000"),
				},
				Error: nil,
			},
		},
		{
			Input: "-000123--4",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("-000123"),
					*lextypes.NewMinus(),
					*lextypes.NewNumber("-4"),
				},
				Error: nil,
			},
		},
	})

	ttng.Run(parser.HandleTestOfParser)
}

func Test__Single_Minus_Signed_Erroring(t *testing.T) {
	ttng := tests.NewTesting(t, []parser.TestOfParser{
		{
			Input: "-0-",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("-0"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(3),
			},
		},
		{
			Input: "-1-",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("-1"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(3),
			},
		},
		{
			Input: "-01-",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("-01"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(4),
			},
		},
		{
			Input: "-000-",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("-000"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(5),
			},
		},
		{
			Input: "-0001-",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("-0001"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(6),
			},
		},
		{
			Input: "-000123-",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("-000123"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(8),
			},
		},
	})

	ttng.Run(parser.HandleTestOfParser)
}

func Test__MultiSigneds_Erroring(t *testing.T) {
	ttng := tests.NewTesting(t, []parser.TestOfParser{
		{
			Input: "0---0",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("0"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(3),
			},
		},
		{
			Input: "-0---0",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("-0"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(4),
			},
		},
		{
			Input: "0-----0",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("0"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(3),
			},
		},
		{
			Input: "+1---1",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("+1"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(4),
			},
		},
		{
			Input: "-1----+2",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("-1"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(4),
			},
		},
		{
			Input: "-12--+34",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("-12"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(5),
			},
		},
		{
			Input: "+56-----78",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("+56"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(5),
			},
		},
		{
			Input: "-123---456",
			Expected: &parser.Expected{
				Tokens: []lextypes.Token{
					*lextypes.NewNumber("-123"),
					*lextypes.NewMinus(),
				},
				Error: lexerrors.NewErrNumberExpected(6),
			},
		},
	})

	ttng.Run(parser.HandleTestOfParser)
}
