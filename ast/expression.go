package ast

import (
	"monkey/token"
	"strconv"
)

type Expression interface {
	Node
	expressionNode()
}

type Identifier struct {
	Token *token.Token // token.IDENT トークン
	Value string
}

var _ Expression = (*Identifier)(nil)

func NewIdentifier(token *token.Token) *Identifier {
	return &Identifier{
		Token: token,
		Value: token.Literal,
	}
}

func NewIdentifierByName(name string) *Identifier {
	return NewIdentifier(token.NewIdentifierToken(name))
}

func (i Identifier) expressionNode() {}

func (i Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i Identifier) String() string {
	return i.Value
}

type IntegerLiteral struct {
	Token *token.Token // token.INT トークン
	Value int64
}

var _ Expression = (*IntegerLiteral)(nil)

func NewIntegerLiteral(token *token.Token, value int64) *IntegerLiteral {
	return &IntegerLiteral{
		Token: token,
		Value: value,
	}
}

func NewIntegerLiteralByValue(value int64) *IntegerLiteral {
	return NewIntegerLiteral(token.NewIntegerToken(strconv.FormatInt(value, 10)), value)
}

func (i IntegerLiteral) expressionNode() {}

func (i IntegerLiteral) TokenLiteral() string {
	return i.Token.Literal
}

func (i IntegerLiteral) String() string {
	return i.Token.Literal
}