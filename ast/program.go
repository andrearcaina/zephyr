package ast

import "bytes"

type Program struct {
	Statements []Statement
}

// TokenLiteral is a recursive function that returns the token literal (the value of the first token)
// of the first statement in the program.
// It then returns an empty string if there are no statements left given the length of the programs statements field
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// String creates a buffer and writes the return value of each statement's String method to it.
// It then returns the string representation of the buffer. It delegates most of the work to the Statements of the Program struct.
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
