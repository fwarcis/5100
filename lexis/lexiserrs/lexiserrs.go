package lexiserrs

import "fmt"

func sprefixf(format string, position int, a ...any) string {
	return fmt.Sprintf("lexis:%d: "+format, position, a)
}

type ErrIllegalChar struct {
	Char     rune
	Position int
}

func (e *ErrIllegalChar) Error() string {
	return sprefixf("illegal char '%s'", e.Position, e.Char)
}

type ErrUnexpectedChar struct {
	Char     rune
	Position int
}

func (e *ErrUnexpectedChar) Error() string {
	return sprefixf("unexpected char '%s'", e.Position, e.Char)
}

type ErrNoMoreTokens struct {
	Position int
}

func (e *ErrNoMoreTokens) Error() string {
	return sprefixf("no more tokens", e.Position)
}
