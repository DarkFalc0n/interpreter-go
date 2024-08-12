package token

import (
	"fmt"
	"os"
)

const (
	NONE = iota
	UNEXPECTED_TOKEN
	UNTERMINATED_STRING
)

var TokenError = 0

func ThrowUnexpectedTokenError (line int, lexeme string) {
	fmt.Fprintln(os.Stderr, "[line " + fmt.Sprint(line) + "] Error: Unexpected character: " + lexeme)
}

func ThrowUnterminatedStringError (line int) {
	fmt.Fprintln(os.Stderr, "[line " + fmt.Sprint(line) + "] Error: Unterminated string.")
}