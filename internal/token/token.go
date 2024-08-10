package token

import "fmt"

type Token struct {
	tokenType int
	literal string
}

func NewToken(tokenType int, literal string) *Token{
	return &Token{
		tokenType: tokenType,
		literal: literal,
	}
}

func (t *Token) Print() {
	fmt.Println(TokenNames[t.tokenType] + " " + TokenLexemes[t.tokenType] + " " + t.literal)
}
