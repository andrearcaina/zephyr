package ast

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
