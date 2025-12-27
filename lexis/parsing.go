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

var numberSigns = []string{"-", "+"}

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
			if numSign == "-" && char == "-" || numSign == "+" && char == "+" {
				numSign = "+"
			} else if numSign == "+" && char == "-" || numSign == "-" && char == "+" {
				numSign = "-"
			} else if unicode.IsDigit(r) {
				digits += char
			}
			state = DigitOrBinOp
		case DigitOrBinOp:
			if unicode.IsDigit(r) {
				digits += char
			} else if slices.Contains(binOpValues, BinOpValue(r)) {
				binOpVal := BinOpValue(r)
				if binOpVal == MinusValue {
					binOpVal = PlusValue
				}
				tokens = append(tokens, *NewNumber(string(numSign) + digits))
				tokens = append(tokens, *binOps[binOpVal]())
				digits = ""
				if char == "-" {
					numSign = "-"
				} else {
					numSign = "+"
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
		!slices.Contains(binOpValues, BinOpValue(r))
}

func illegal(r rune, state parsingState) bool {
	switch state {
	case SignOrDigit:
		if !slices.Contains(numberSigns, string(r)) && !unicode.IsDigit(r) {
			return true
		}
	case DigitOrBinOp:
		if !slices.Contains(binOpValues, BinOpValue(r)) && !unicode.IsDigit(r) {
			return true
		}
	}
	return false
}
