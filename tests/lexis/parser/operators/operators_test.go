package operators_test

import (
	"testing"

	"5100/lexis/lexerrors"
	"5100/tests"
	"5100/tests/lexis/parser"
)

func run(t *testing.T, opVal string, posOfUnexpect int) {
	inputs := []string{
		opVal,
		opVal + opVal,
		opVal + opVal + opVal,
	}

	cases := make([]parser.TestOfParser, 0, len(inputs))
	for _, inp := range inputs {
		cases = append(cases, parser.TestOfParser{
			Input: inp,
			Expected: &parser.Expected{
				Error: lexerrors.NewErrNumberExpected(posOfUnexpect),
			},
		})
	}
	ttng := tests.NewTesting(t, cases)

	ttng.Run(parser.HandleTestOfParser)
}

func Test__Plus(t *testing.T) {
	run(t, "+", 1)
}

func Test__Minus(t *testing.T) {
	run(t, "-", 1)
}

func Test__Multiplication(t *testing.T) {
	run(t, "*", 0)
}

func Test__Division(t *testing.T) {
	run(t, "/", 0)
}

func Test__Jointly(t *testing.T) {
	inputs := []string{
		"+-*/",
		"*/-",
		"/*-+*",
		"-+*",
		"*/+-+",
		"+-*",
		"-/*+",
		"+*/",
		"*+-/",
		"/-*+",
	}
	cases := make([]parser.TestOfParser, 0, len(inputs))
	for _, inp := range inputs {
		var charPosForErr int
		if inp[0] == '+' || inp[0] == '-' {
			charPosForErr = 1
		} else {
			charPosForErr = 0
		}
		cases = append(cases, parser.TestOfParser{
			Input: inp,
			Expected: &parser.Expected{
				Error: lexerrors.NewErrNumberExpected(charPosForErr),
			},
		})
	}
	ttng := tests.NewTesting(t, cases)

	ttng.Run(parser.HandleTestOfParser)
}
