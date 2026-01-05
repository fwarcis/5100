package lextypes

func NewNumber(value string) *Token {
	return &Token{
		Value: value,
		Type:  NumberType,
	}
}
