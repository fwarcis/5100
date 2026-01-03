package lextypes

func NewNumber(value string) *Token {
	return &Token{
		Value:    value,
		Priority: 100,
		Type:     NumberType,
	}
}
