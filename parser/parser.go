package parser

import (
	"github.com/andrearcaina/zephyr/ast"
	"github.com/andrearcaina/zephyr/lexer"
	"github.com/andrearcaina/zephyr/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string

	// prefixParseFns is a map of functions, where the token Category is the key,
	// and the value is a function that returns an ast.Expression (e.g., an identifier, integer, or boolean).
	prefixParseFns map[token.Category]func() ast.Expression
	// infixParseFns is a map of functions, where the token Category is the key,
	// and the value is a function that takes an ast.Expression and returns another ast.Expression (e.g., for parsing binary operations).
	// an example of this is the addition operator, which takes an ast.Expression (the left operand) and returns a new ast.Expression (the result of the addition).
	infixParseFns map[token.Category]func(ast.Expression) ast.Expression
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	p.prefixParseFns = make(map[token.Category]func() ast.Expression)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)

	return p
}
