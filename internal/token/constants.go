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
	"EQUAL",
	"EQUAL_EQUAL",
	"BANG",
	"BANG_EQUAL",
	"LESS",
	"LESS_EQUAL",
	"GREATER",
	"GREATER_EQUAL",
	"SLASH",
	"STRING",
	"NUMBER",
	"IDENTIFIER",
	"EOF",
	"INVALID",
}


var reservedKeywords = []string {
	AND : "and",
	CLASS : "class",
	ELSE : "else",
	FALSE : "false",
	FOR : "for",
	FUN : "fun",
	IF : "if",
	NIL : "nil",
	OR : "or",
	PRINT : "print",
	RETURN : "return",
	SUPER : "super",
	THIS : "this",
	TRUE : "true",
	VAR : "var",
	WHILE : "while",
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
	EQUAL
	EQUAL_EQUAL
	BANG
	BANG_EQUAL
	LESS
	LESS_EQUAL
	GREATER
	GREATER_EQUAL
	SLASH
	STRING
	NUMBER
	IDENTIFIER	
	EOF
	INVALID
	// reserved keywords
	AND
	CLASS
	ELSE
	FALSE
	FOR
	FUN
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
)