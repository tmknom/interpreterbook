package parser

import (
	"fmt"
	"github.com/pkg/errors"
	"monkey/ast"
	"monkey/token"
	"strconv"
)

func (p *Parser) parseExpression(precedence precedence) ast.Expression {
	prefix := p.prefixParseFns[p.currentToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.currentToken.Type)
		return nil
	}

	leftExp := prefix()
	return leftExp
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	message := fmt.Sprintf("no prefix parse function for %q found", t)
	err := errors.New(message)
	p.errors = append(p.errors, err)
}

func (p *Parser) parseIdentifier() ast.Expression {
	return ast.NewIdentifier(p.currentToken)
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	value, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)
	if err != nil {
		message := fmt.Sprintf("could not parse %q as integer", p.peekToken.Literal)
		err := errors.New(message)
		p.errors = append(p.errors, err)
		return nil
	}

	return ast.NewIntegerLiteral(p.currentToken, value)
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := ast.NewPrefixExpression(p.currentToken)

	// 前置トークンの次を参照するため、ひとつ進めておく
	p.nextToken()

	right := p.parseExpression(PREFIX)
	expression.SetRight(right)

	return expression
}

func (p *Parser) initExpressionFunctions() {
	p.prefixParseFns = map[token.TokenType]prefixParseFn{}
	p.infixParseFns = map[token.TokenType]infixParseFn{}

	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}
