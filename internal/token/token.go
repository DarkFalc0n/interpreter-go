package token

import (
	"fmt"
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
		ThrowUnexpectedTokenError(t.line, t.lexeme)
		return
	}
	fmt.Println(TokenNames[t.tokenType] + " " + t.lexeme + " " + string(t.literal))
}
