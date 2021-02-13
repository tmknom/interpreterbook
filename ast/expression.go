package ast

import (
	"bytes"
	"monkey/token"
	"strconv"
	"strings"
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

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
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

func (i *IntegerLiteral) expressionNode() {}

func (i *IntegerLiteral) TokenLiteral() string {
	return i.Token.Literal
}

func (i *IntegerLiteral) String() string {
	return i.Token.Literal
}

type Boolean struct {
	Token *token.Token
	Value bool
}

var _ Expression = (*Boolean)(nil)

func NewBoolean(token *token.Token, value bool) *Boolean {
	return &Boolean{
		Token: token,
		Value: value,
	}
}

func NewBooleanByValue(value string) *Boolean {
	return NewBoolean(token.NewIdentifierToken(value), value == "true")
}

func (b *Boolean) expressionNode() {}

func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Token.Literal
}

type PrefixExpression struct {
	Token    *token.Token // 前置トークン／たとえば「!」
	Operator string
	Right    Expression
}

var _ Expression = (*PrefixExpression)(nil)

func NewPrefixExpression(token *token.Token) *PrefixExpression {
	return &PrefixExpression{
		Token:    token,
		Operator: token.Literal,
	}
}

func (e *PrefixExpression) SetRight(right Expression) {
	e.Right = right
}

func (e *PrefixExpression) expressionNode() {}

func (e *PrefixExpression) TokenLiteral() string {
	return e.Token.Literal
}

func (e *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(e.Operator)
	out.WriteString(e.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    *token.Token // 演算子トークン／たとえば「+」
	Left     Expression
	Operator string
	Right    Expression
}

var _ Expression = (*InfixExpression)(nil)

func NewInfixExpression(token *token.Token, left Expression) *InfixExpression {
	return &InfixExpression{
		Token:    token,
		Operator: token.Literal,
		Left:     left,
	}
}

func (e *InfixExpression) SetRight(right Expression) {
	e.Right = right
}

func (e *InfixExpression) expressionNode() {}

func (e *InfixExpression) TokenLiteral() string {
	return e.Token.Literal
}

func (e *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(e.Left.String())
	out.WriteString(" " + e.Operator + " ")
	out.WriteString(e.Right.String())
	out.WriteString(")")

	return out.String()
}

type IfExpression struct {
	Token       *token.Token // 'if' トークン
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

var _ Expression = (*IfExpression)(nil)

func NewIfExpression(token *token.Token) *IfExpression {
	return &IfExpression{
		Token: token,
	}
}

func (e *IfExpression) SetCondition(exp Expression) {
	e.Condition = exp
}

func (e *IfExpression) SetConsequence(bs *BlockStatement) {
	e.Consequence = bs
}

func (e *IfExpression) SetAlternative(bs *BlockStatement) {
	e.Alternative = bs
}

func (e *IfExpression) expressionNode() {}

func (e *IfExpression) TokenLiteral() string {
	return e.Token.Literal
}

func (e *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(e.Condition.String())
	out.WriteString(" ")
	out.WriteString(e.Consequence.String())

	if e.Alternative != nil {
		out.WriteString("else")
		out.WriteString(e.Alternative.String())
	}

	return out.String()
}

type FunctionLiteral struct {
	Token      *token.Token // 'fn' トークン
	Parameters []*Identifier
	Body       *BlockStatement
}

var _ Expression = (*FunctionLiteral)(nil)

func NewFunctionLiteral(token *token.Token) *FunctionLiteral {
	return &FunctionLiteral{
		Token:      token,
		Parameters: []*Identifier{},
	}
}

func (l *FunctionLiteral) SetParameters(parameters []*Identifier) {
	l.Parameters = parameters
}

func (l *FunctionLiteral) SetBody(body *BlockStatement) {
	l.Body = body
}

func (l *FunctionLiteral) expressionNode() {}

func (l *FunctionLiteral) TokenLiteral() string {
	return l.Token.Literal
}

func (l *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, parameter := range l.Parameters {
		params = append(params, parameter.String())
	}

	out.WriteString(l.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(")")
	out.WriteString(l.Body.String())

	return out.String()
}
