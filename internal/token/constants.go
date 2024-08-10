package token

var TokenNames = []string {
	"LEFT_PAREN",
	"RIGHT_PAREN",
	"EOF",
}

var TokenLexemes = []string {
	"(",
	")",
	"",
}

const (
	LEFT_PAREN = iota
	RIGHT_PAREN
	EOF
)