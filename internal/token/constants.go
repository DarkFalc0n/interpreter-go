package token

var TokenNames = []string {
	"LEFT_PAREN",
	"RIGHT_PAREN",
	"LEFT_BRACE",
	"RIGHT_BRACE",
	"COMMA",
	"DOT",
	"MINUS",
	"PLUS",
	"SEMICOLON",
	"STAR",
	"EOF",
}

var TokenLexemes = []string {
	"(",
	")",
	"{",
	"}",
	",",
	".",
	"-",
	"+",
	";",
	"*",
	"",
}

const (
	LEFT_PAREN = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	STAR
	EOF
)