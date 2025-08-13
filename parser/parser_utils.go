package parser

import (
	"fmt"

	"github.com/andrearcaina/zephyr/ast"
	"github.com/andrearcaina/zephyr/token"
)

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.Category) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.Category) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.Category) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.Category) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) registerPrefix(tokenType token.Category, fn func() ast.Expression) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.Category, fn func(ast.Expression) ast.Expression) {
	p.infixParseFns[tokenType] = fn
}
