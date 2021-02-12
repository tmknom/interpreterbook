package parser

import "monkey/ast"

func (p *Parser) ParseProgram() *ast.Program {
	program := ast.NewProgram()
	for !p.currentToken.IsEOF() {
		stmt := p.parseStatement()
		if stmt != nil {
			program.AddStatement(stmt)
		}
		p.nextToken()
	}
	return program
}
