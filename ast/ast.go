package ast

import (
	"fmt"
	"monkey/token"
)

type Node interface {
	TokenLiteral() string
	fmt.Stringer
}

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
	}
}

func NewIdentifierByName(name string) *Identifier {
	return &Identifier{
		Token: token.NewIdentifierToken(name),
	}
}

func (i Identifier) expressionNode() {}

func (i Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i Identifier) String() string {
	return i.Value
}
