package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

var _ Node = (*Program)(nil)

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token *token.Token // token.LET トークン
	Name  *Identifier
	Value Expression
}

var _ Statement = (*LetStatement)(nil)

func (s LetStatement) statementNode() {}

func (s LetStatement) TokenLiteral() string {
	return s.Token.Literal
}

type Identifier struct {
	Token *token.Token // token.IDENT トークン
	Value Expression
}

var _ Expression = (*Identifier)(nil)

func (i Identifier) expressionNode() {}

func (i Identifier) TokenLiteral() string {
	return i.Token.Literal
}
