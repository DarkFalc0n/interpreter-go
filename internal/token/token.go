package token

import (
	"fmt"
	"os"
)

type Token struct {
	tokenType int
	lexeme string
	literal string
	line int
}

func NewToken(tokenType int, lexeme string, literal string, line int) *Token{
	return &Token{
		tokenType: tokenType,
		lexeme: lexeme,
		literal: literal,
		line: line,
	}
}

func (t *Token) Print() {
	if t.tokenType == INVALID {
		HasError = true
		fmt.Fprintf(os.Stderr, "[line " + fmt.Sprint(t.line) + "] Error: Invalid character: " + t.lexeme + "\n")
		return
	}
	fmt.Println(TokenNames[t.tokenType] + " " + t.lexeme + " " + t.literal)
}
