package token

import (
	"fmt"
	"strings"
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
	if t.tokenType >= AND && t.tokenType <= WHILE {
		fmt.Println(strings.ToUpper(reservedKeywords[t.tokenType]) + " " + t.lexeme + " " + string(t.literal))
		return
	}
	fmt.Println(TokenNames[t.tokenType] + " " + t.lexeme + " " + string(t.literal))
}
