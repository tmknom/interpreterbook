package parser

import (
	"monkey/ast"
	"monkey/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	if !p.currentTokenIs(token.LET) {
		return nil
	}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	name := ast.NewIdentifier(p.currentToken)
	stmt := ast.NewLetStatement(name)

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO セミコロンに遭遇するまで読み飛ばす
	for !p.currentTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	if !p.currentTokenIs(token.RETURN) {
		return nil
	}

	p.nextToken()

	stmt := ast.NewReturnStatement()

	// TODO セミコロンに遭遇するまで読み飛ばす
	for !p.currentTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
