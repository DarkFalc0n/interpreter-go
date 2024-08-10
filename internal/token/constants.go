package token

var TokenNames = []string {
	"LEFT_PAREN",
	"RIGHT_PAREN",
	"LEFT_BRACE",
	"RIGHT_BRACE",
	"EOF",
}

var TokenLexemes = []string {
	"(",
	")",
	"{",
	"}",
	"",
}

const (
	LEFT_PAREN = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	EOF
)