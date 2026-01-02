package numbers_test

import (
	"testing"

	"5100/lexis"
	"5100/tests"
	"5100/tests/lexis/parser"
)

func Test__NoSign(t *testing.T) {
	ttng := tests.NewTesting(t, []parser.TestOfParser{
		{
			Input: "0",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("0"),
				},
				Error: nil,
			},
		},
		{
			Input: "000",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("000"),
				},
				Error: nil,
			},
		},
		{
			Input: "1",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("1"),
				},
				Error: nil,
			},
		},
		{
			Input: "123",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("123"),
				},
				Error: nil,
			},
		},
		{
			Input: "123000",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("123000"),
				},
				Error: nil,
			},
		},
		{
			Input: "01",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("01"),
				},
				Error: nil,
			},
		},
		{
			Input: "0001",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("0001"),
				},
				Error: nil,
			},
		},
		{
			Input: "000123",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("000123"),
				},
				Error: nil,
			},
		},
		{
			Input: "000123000",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("000123000"),
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
			Input: "+0",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+0"),
				},
				Error: nil,
			},
		},
		{
			Input: "+000",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+000"),
				},
				Error: nil,
			},
		},
		{
			Input: "+1",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+1"),
				},
				Error: nil,
			},
		},
		{
			Input: "+123",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+123"),
				},
				Error: nil,
			},
		},
		{
			Input: "+123000",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+123000"),
				},
				Error: nil,
			},
		},
		{
			Input: "+01",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+01"),
				},
				Error: nil,
			},
		},
		{
			Input: "+0001",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+0001"),
				},
				Error: nil,
			},
		},
		{
			Input: "+000123",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+000123"),
				},
				Error: nil,
			},
		},
		{
			Input: "+000123000",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("+000123000"),
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
			Input: "-0",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-0"),
				},
				Error: nil,
			},
		},
		{
			Input: "-000",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-000"),
				},
				Error: nil,
			},
		},
		{
			Input: "-1",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-1"),
				},
				Error: nil,
			},
		},
		{
			Input: "-123",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-123"),
				},
				Error: nil,
			},
		},
		{
			Input: "-123000",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-123000"),
				},
				Error: nil,
			},
		},
		{
			Input: "-01",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-01"),
				},
				Error: nil,
			},
		},
		{
			Input: "-0001",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-0001"),
				},
				Error: nil,
			},
		},
		{
			Input: "-000123",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-000123"),
				},
				Error: nil,
			},
		},
		{
			Input: "-000123000",
			Expected: &parser.Expected{
				Tokens: []lexis.Token{
					*lexis.NewNumber("-000123000"),
				},
				Error: nil,
			},
		},
	})

	ttng.Run(parser.HandleTestOfParser)
}

func Test__MultiSigneds_Erroring(t *testing.T) {
	inputs := []string{
		"+-0",
		"+-000",
		"+-1",
		"+-123",
		"+-123000",
		"+-01",
		"+-0001",
		"+-000123",
		"+-000123000",

		"-+0",
		"-+000",
		"-+1",
		"-+123",
		"-+123000",
		"-+01",
		"-+0001",
		"-+000123",
		"-+000123000",

		"++-0",
		"--+0",
		"+-+-1",
		"-+-+1",
		"+++---123",
		"---+++123",
		"+-+-+-0001",
		"-+-+-+0001",
		"++--+-000123",
	}
	cases := make([]parser.TestOfParser, 0, len(inputs))
	for _, inp := range inputs {
		cases = append(cases, parser.TestOfParser{
			Input: inp,
			Expected: &parser.Expected{
				Error: lexis.NewErrNumberExpected(1),
			},
		})
	}
	ttng := tests.NewTesting(t, cases)

	ttng.Run(parser.HandleTestOfParser)
}
