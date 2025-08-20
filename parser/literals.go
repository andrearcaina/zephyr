package parser

import (
	"strconv"

	"github.com/andrearcaina/zephyr/ast"
)

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		p.errors = append(p.errors, "could not parse integer literal: "+p.curToken.Literal)
		return nil
	}

	lit.Value = value
	return lit
}
