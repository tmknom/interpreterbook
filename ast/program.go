package ast

import "bytes"

type Program struct {
	Statements []Statement
}

var _ Node = (*Program)(nil)

func NewProgram() *Program {
	return &Program{
		Statements: []Statement{},
	}
}

func (p *Program) AddStatement(stmt Statement) {
	p.Statements = append(p.Statements, stmt)
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, stmt := range p.Statements {
		out.WriteString(stmt.String())
	}
	return out.String()
}
