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
		return nil
	}

	leftExp := prefix()
	return leftExp
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

func (p *Parser) initExpressionFunctions() {
	p.prefixParseFns = map[token.TokenType]prefixParseFn{}
	p.infixParseFns = map[token.TokenType]infixParseFn{}

	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}
