package parser

import (
	"fmt"
	"github.com/pkg/errors"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type prefixParseFn func() ast.Expression              // 前置構文解析関数
type infixParseFn func(ast.Expression) ast.Expression // 中置構文解析関数

type Parser struct {
	l            *lexer.Lexer
	currentToken *token.Token
	peekToken    *token.Token
	errors       []error

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []error{},
	}

	// 2つトークンを読み込む
	// currentTokenとpeekTokenの両方がセットされる
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) currentTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) peekError(t token.TokenType) {
	message := fmt.Sprintf("expected next token to be '%s', got: '%s', detail: %s", t, p.peekToken.Type, p.peekToken.Detail())
	err := errors.New(message)
	p.errors = append(p.errors, err)
}

func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}
