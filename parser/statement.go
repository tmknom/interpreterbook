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
		return p.parseExpressionStatement()
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

	p.nextToken()

	expression := p.parseExpression(LOWEST)
	stmt.SetValue(expression)

	if p.peekTokenIs(token.SEMICOLON) {
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

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := ast.NewExpressionStatement(p.currentToken)
	exp := p.parseExpression(LOWEST)
	stmt.SetExpression(exp)

	// セミコロンを省略可能にするため、セミコロンを見つけたらひとつ進める
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	blockStatement := ast.NewBlockStatement(p.currentToken)
	p.nextToken()

	for !p.currentTokenIs(token.RBRACE) && !p.currentTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			blockStatement.AddStatement(stmt)
		}
		p.nextToken()
	}

	return blockStatement
}
