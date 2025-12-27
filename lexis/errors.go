package lexis

import "fmt"

type ErrIllegalChar struct {
	char     string
	position int
}

func (err *ErrIllegalChar) Error() string {
	return fmt.Sprintf("lexis:%d: illegal char '%s'", err.position, err.char)
}

type ErrUnexpectedChar struct {
	char     string
	position int
}

func (err *ErrUnexpectedChar) Error() string {
	return fmt.Sprintf("lexis:%d: unexpected char '%s'", err.position, err.char)
}
