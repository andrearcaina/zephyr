package token

// Token represents a single token in the input given to the lexer

type Category string

type Token struct {
	Type    Category
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	BANG     = "!"

	// Comparison operators
	GT  = ">"
	LT  = "<"
	EQ  = "=="
	NEQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LBRACE = "{"
	RBRACE = "}"
	LPAREN = "("
	RPAREN = ")"

	// Keywords
	FUNCTION = "FUNC"
	LET      = "LET"
	CONST    = "CONST"
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
)

var keywords = map[string]Category{
	"func":   FUNCTION,
	"let":    LET,
	"const":  CONST,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
}

// LookupIdent checks if the given identifier is a keyword and returns the corresponding token type
func LookupIdentifier(identifier string) Category {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENT
}
