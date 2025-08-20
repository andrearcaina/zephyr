package ast

import (
	"bytes"

	"github.com/andrearcaina/zephyr/token"
)

type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

func (s *LetStatement) statementNode()       {}
func (s *LetStatement) TokenLiteral() string { return s.Token.Literal }
func (s *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(s.TokenLiteral() + " ")
	out.WriteString(s.Name.String())
	out.WriteString(" = ")

	if s.Value != nil {
		out.WriteString(s.Value.String())
	} else {
		out.WriteString("nil")
	}

	out.WriteString(";")
	return out.String()
}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (s *ReturnStatement) statementNode()       {}
func (s *ReturnStatement) TokenLiteral() string { return s.Token.Literal }
func (s *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(s.TokenLiteral() + " ")

	if s.Value != nil {
		out.WriteString(s.Value.String())
	} else {
		out.WriteString("nil")
	}

	out.WriteString(";")
	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (s *ExpressionStatement) statementNode()       {}
func (s *ExpressionStatement) TokenLiteral() string { return s.Token.Literal }
func (s *ExpressionStatement) String() string {
	if s.Expression != nil {
		return s.Expression.String()
	}
	return ""
}
