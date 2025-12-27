package lexis

import (
	"slices"
	"unicode"
)

const (
	SignOrDigit = iota
	DigitOrBinOp
)

var numSigns = []rune{'+', '-'}

func Parse(text string) (*[]Token, error) {
	runes := []rune(text)
	tokens := []Token{}
	digits := ""
	numSign := "+"
	state := SignOrDigit
	for i, r := range runes {
		char := string(r)

		if illegal(r, state) {
			return nil, &ErrIllegalChar{char: char, position: i}
		} else if unexpected(r) {
			return nil, &ErrUnexpectedChar{char: char, position: i}
		}

		switch state {
		case SignOrDigit:
			if numSign == "-" && r == '-' || numSign == "+" && r == '+' {
				numSign = "+"
			} else if numSign == "+" && r == '-' || numSign == "-" && r == '+' {
				numSign = "-"
			} else if unicode.IsDigit(r) {
				digits += string(r)
			}
			state = DigitOrBinOp
		case DigitOrBinOp:
			if unicode.IsDigit(r) {
				digits += string(r)
			} else if slices.Contains(binOpValues, string(r)) {
				binOpVal := string(r)
				if binOpVal == "-" {
					binOpVal = "+"
				}
				tokens = append(tokens, *NewNumber(numSign + digits))
				tokens = append(tokens, *NewBinaryOperator(binOpVal))
				digits = ""
				if string(r) == "-" {
					numSign = "-"
				} else {
					numSign = "+"
				}
				state = SignOrDigit
			}
		}

		if i+1 == len(runes) {
			tokens = append(tokens, *NewNumber(numSign + digits))
		}
	}
	return &tokens, nil
}

func unexpected(r rune) bool {
	return !unicode.IsDigit(r) &&
		!slices.Contains(binOpValues, string(r))
}

func illegal(r rune, state int) bool {
	switch state {
	case SignOrDigit:
		if !slices.Contains(numSigns, r) && !unicode.IsDigit(r) {
			return true
		}
	case DigitOrBinOp:
		if !slices.Contains(binOpValues, string(r)) && !unicode.IsDigit(r) {
			return true
		}
	}
	return false
}
