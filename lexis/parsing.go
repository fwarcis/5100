package lexis

import (
	"slices"
	"unicode"
)

type parsingState int

const (
	SignOrDigit parsingState = iota
	DigitOrBinOp
)

type numberSign string

const (
	minusSign numberSign = "-"
	plusSign  numberSign = "+"
)

var numberSigns = []numberSign{minusSign, plusSign}

func Parse(text string) (*[]Token, error) {
	runes := []rune(text)
	tokens := []Token{}
	digits := ""
	numSign := plusSign
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
			if numSign == minusSign && r == '-' || numSign == plusSign && r == '+' {
				numSign = plusSign
			} else if numSign == plusSign && r == '-' || numSign == minusSign && r == '+' {
				numSign = minusSign
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
				tokens = append(tokens, *NewNumber(string(numSign) + digits))
				tokens = append(tokens, *binOps[binOpVal]())
				digits = ""
				if string(r) == "-" {
					numSign = minusSign
				} else {
					numSign = plusSign
				}
				state = SignOrDigit
			}
		}

		if i+1 == len(runes) {
			tokens = append(tokens, *NewNumber(string(numSign) + digits))
		}
	}
	return &tokens, nil
}

func unexpected(r rune) bool {
	return !unicode.IsDigit(r) &&
		!slices.Contains(binOpValues, string(r))
}

func illegal(r rune, state parsingState) bool {
	switch state {
	case SignOrDigit:
		if !slices.Contains(numberSigns, numberSign(r)) && !unicode.IsDigit(r) {
			return true
		}
	case DigitOrBinOp:
		if !slices.Contains(binOpValues, string(r)) && !unicode.IsDigit(r) {
			return true
		}
	}
	return false
}
