package parser

import (
	"fmt"
	"github.com/pkg/errors"
	"monkey/ast"
	"monkey/token"
	"strconv"
)

func (p *Parser) parseExpression(precedence precedence) ast.Expression {
	trace(fmt.Sprintf("parseExpression(): {%s}", p.debug()))

	prefix := p.prefixParseFns[p.currentToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.currentToken.Type)
		return nil
	}
	leftExp := prefix()

	traceDetail(fmt.Sprintf("if precedence(%d) < p.peekPrecedence(%d) then call infixParseFn", precedence, p.peekPrecedence()))
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
		traceDetail(fmt.Sprintf("if precedence(%d) < p.peekPrecedence(%d) then call infixParseFn", precedence, p.peekPrecedence()))
	}

	untrace(fmt.Sprintf("parseExpression() => return Expression{%q}", leftExp))
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
	trace(fmt.Sprintf("parseIntegerLiteral(): {%s}", p.debug()))

	value, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)
	if err != nil {
		message := fmt.Sprintf("could not parse %q as integer", p.peekToken.Literal)
		err := errors.New(message)
		p.errors = append(p.errors, err)
		return nil
	}

	expression := ast.NewIntegerLiteral(p.currentToken, value)
	untrace(fmt.Sprintf("parseIntegerLiteral() => return IntegerLiteral{%q}", expression))
	return expression
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	trace(fmt.Sprintf("parsePrefixExpression(): {%s}", p.debug()))

	expression := ast.NewPrefixExpression(p.currentToken)

	// 前置トークンの次を参照するため、ひとつ進めておく
	p.nextToken()

	right := p.parseExpression(PREFIX)
	expression.SetRight(right)

	untrace(fmt.Sprintf("parsePrefixExpression() => return PrefixExpression{%q}", expression))
	return expression
}

func (p *Parser) parseBoolean() ast.Expression {
	trace(fmt.Sprintf("parseBoolean(): {%s}", p.debug()))

	boolean := ast.NewBoolean(p.currentToken, p.currentTokenIs(token.TRUE))

	untrace(fmt.Sprintf("parseBoolean() => return Boolean{%q}", boolean))
	return boolean
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	trace(fmt.Sprintf("parseInfixExpression(left=%q): {%s}", left, p.debug()))

	traceDetail(fmt.Sprintf("new InfixExpression{left=%q, operator=%q}", left, p.currentToken.Literal))
	expression := ast.NewInfixExpression(p.currentToken, left)

	precedence := p.currentPrecedence()
	p.nextToken()

	right := p.parseExpression(precedence)
	expression.SetRight(right)

	untrace(fmt.Sprintf("parseInfixExpression() => return InfixExpression{%q}", expression))
	return expression
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	trace(fmt.Sprintf("parseGroupedExpression(): {%s}", p.debug()))

	p.nextToken()
	expression := p.parseExpression(LOWEST)

	traceDetail(fmt.Sprintf("if p.expectPeek(token.RPAREN) then ok: {%s}", p.debug()))
	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	untrace(fmt.Sprintf("parseGroupedExpression() => return Expression{%q}", expression))
	return expression
}

func (p *Parser) initExpressionFunctions() {
	p.prefixParseFns = map[token.TokenType]prefixParseFn{}
	p.infixParseFns = map[token.TokenType]infixParseFn{}

	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)

	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)

	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)

	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}
