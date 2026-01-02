package plus_test

import (
	"testing"

	"5100/lexis"
	"5100/tests"
	"5100/tests/lexis/parser"
)

func Test__NoSign(t *testing.T) {
	ttng := tests.NewTesting(t, []parser.TestOfParser{
		{
			Input: "0+0",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("0"),
					*lexis.NewPlus(),
					*lexis.NewNumber("0"),
				},
				Error: nil,
			},
		},
		{
			Input: "0+1",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("0"),
					*lexis.NewPlus(),
					*lexis.NewNumber("1"),
				},
				Error: nil,
			},
		},
		{
			Input: "1+0",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("1"),
					*lexis.NewPlus(),
					*lexis.NewNumber("0"),
				},
				Error: nil,
			},
		},
		{
			Input: "01+1",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("01"),
					*lexis.NewPlus(),
					*lexis.NewNumber("1"),
				},
				Error: nil,
			},
		},
		{
			Input: "1+01",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("1"),
					*lexis.NewPlus(),
					*lexis.NewNumber("01"),
				},
				Error: nil,
			},
		},
		{
			Input: "000+000",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("000"),
					*lexis.NewPlus(),
					*lexis.NewNumber("000"),
				},
				Error: nil,
			},
		},
		{
			Input: "0001+002",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("0001"),
					*lexis.NewPlus(),
					*lexis.NewNumber("002"),
				},
				Error: nil,
			},
		},
	})

	ttng.Run(parser.HandleTestOfParser)
}

func Test__Positive_Signeds(t *testing.T) {
	ttng := tests.NewTesting(t, []parser.TestOfParser{
		{
			Input: "+0+0",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+0"),
					*lexis.NewPlus(),
					*lexis.NewNumber("0"),
				},
				Error: nil,
			},
		},
		{
			Input: "+1+0",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+1"),
					*lexis.NewPlus(),
					*lexis.NewNumber("0"),
				},
				Error: nil,
			},
		},
		{
			Input: "+0+1",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+0"),
					*lexis.NewPlus(),
					*lexis.NewNumber("1"),
				},
				Error: nil,
			},
		},
		{
			Input: "+01+1",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+01"),
					*lexis.NewPlus(),
					*lexis.NewNumber("1"),
				},
				Error: nil,
			},
		},
		{
			Input: "+1+01",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+1"),
					*lexis.NewPlus(),
					*lexis.NewNumber("01"),
				},
				Error: nil,
			},
		},
		{
			Input: "+000+000",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+000"),
					*lexis.NewPlus(),
					*lexis.NewNumber("000"),
				},
				Error: nil,
			},
		},
		{
			Input: "+9+0000",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+9"),
					*lexis.NewPlus(),
					*lexis.NewNumber("0000"),
				},
				Error: nil,
			},
		},
		{
			Input: "+000123+4",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+000123"),
					*lexis.NewPlus(),
					*lexis.NewNumber("4"),
				},
				Error: nil,
			},
		},
	})

	ttng.Run(parser.HandleTestOfParser)
}

func Test__Negative_Signeds(t *testing.T) {
	ttng := tests.NewTesting(t, []parser.TestOfParser{
		{
			Input: "-0+-0",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-0"),
					*lexis.NewPlus(),
					*lexis.NewNumber("-0"),
				},
				Error: nil,
			},
		},
		{
			Input: "-1+-0",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-1"),
					*lexis.NewPlus(),
					*lexis.NewNumber("-0"),
				},
				Error: nil,
			},
		},
		{
			Input: "-0+-1",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-0"),
					*lexis.NewPlus(),
					*lexis.NewNumber("-1"),
				},
				Error: nil,
			},
		},
		{
			Input: "-01+-1",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-01"),
					*lexis.NewPlus(),
					*lexis.NewNumber("-1"),
				},
				Error: nil,
			},
		},
		{
			Input: "-1+-01",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-1"),
					*lexis.NewPlus(),
					*lexis.NewNumber("-01"),
				},
				Error: nil,
			},
		},
		{
			Input: "-000+-000",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-000"),
					*lexis.NewPlus(),
					*lexis.NewNumber("-000"),
				},
				Error: nil,
			},
		},
		{
			Input: "-9+-0000",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-9"),
					*lexis.NewPlus(),
					*lexis.NewNumber("-0000"),
				},
				Error: nil,
			},
		},
		{
			Input: "-000123+-4",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-000123"),
					*lexis.NewPlus(),
					*lexis.NewNumber("-4"),
				},
				Error: nil,
			},
		},
	})

	ttng.Run(parser.HandleTestOfParser)
}

func Test__MultiSigneds_Erroring(t *testing.T) {
	ttng := tests.NewTesting(t, []parser.TestOfParser{
		{
			Input: "0+++0",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("0"),
					*lexis.NewPlus(),
				},
				Error: lexis.NewErrNumberExpected(3),
			},
		},
		{
			Input: "-0+++0",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-0"),
					*lexis.NewPlus(),
				},
				Error: lexis.NewErrNumberExpected(4),
			},
		},
		{
			Input: "0++++-0",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("0"),
					*lexis.NewPlus(),
				},
				Error: lexis.NewErrNumberExpected(3),
			},
		},
		{
			Input: "+1+++1",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+1"),
					*lexis.NewPlus(),
				},
				Error: lexis.NewErrNumberExpected(4),
			},
		},
		{
			Input: "-1+++++2",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-1"),
					*lexis.NewPlus(),
				},
				Error: lexis.NewErrNumberExpected(4),
			},
		},
		{
			Input: "-12+++34",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-12"),
					*lexis.NewPlus(),
				},
				Error: lexis.NewErrNumberExpected(5),
			},
		},
		{
			Input: "+56++++-78",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+56"),
					*lexis.NewPlus(),
				},
				Error: lexis.NewErrNumberExpected(5),
			},
		},
		{
			Input: "-123+++456",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-123"),
					*lexis.NewPlus(),
				},
				Error: lexis.NewErrNumberExpected(6),
			},
		},
	})

	ttng.Run(parser.HandleTestOfParser)
}
