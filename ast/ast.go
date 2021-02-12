package ast

import (
	"bytes"
	"fmt"
	"monkey/token"
)

type Node interface {
	TokenLiteral() string
	fmt.Stringer
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

func NewProgram() *Program {
	return &Program{
		Statements: []Statement{},
	}
}

func (p *Program) AddStatement(stmt Statement) {
	p.Statements = append(p.Statements, stmt)
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, stmt := range p.Statements {
		out.WriteString(stmt.String())
	}
	return out.String()
}

type LetStatement struct {
	*token.Token // token.LET トークン
	Name         *Identifier
	Value        Expression
}

var _ Statement = (*LetStatement)(nil)

func NewLetStatement(name *Identifier) *LetStatement {
	return &LetStatement{
		Token: letToken,
		Name:  name,
	}
}

func NewLetStatementByName(name string) *LetStatement {
	return NewLetStatement(NewIdentifierByName(name))
}

var letToken = token.NewToken(token.LET, "let")

func (s LetStatement) statementNode() {}

func (s LetStatement) TokenLiteral() string {
	return s.Token.Literal
}

func (s LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(s.TokenLiteral() + " ")
	out.WriteString(s.Name.String())
	out.WriteString(" = ")

	if s.Value != nil {
		out.WriteString(s.Value.String())
	}
	out.WriteString(";")

	return out.String()
}

type ReturnStatement struct {
	*token.Token // token.RETURN トークン
	ReturnValue  Expression
}

var _ Statement = (*ReturnStatement)(nil)

func NewReturnStatement() *ReturnStatement {
	return &ReturnStatement{
		Token: returnToken,
		//ReturnValue: returnValue,
	}
}

var returnToken = token.NewToken(token.RETURN, "return")

func (s ReturnStatement) statementNode() {}

func (s ReturnStatement) TokenLiteral() string {
	return s.Token.Literal
}

func (s ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(s.TokenLiteral() + " ")

	if s.ReturnValue != nil {
		out.WriteString(s.ReturnValue.String())
	}
	out.WriteString(";")

	return out.String()
}

type ExpressionStatement struct {
	*token.Token // 式の最初のトークン
	Expression
}

var _ Statement = (*ExpressionStatement)(nil)

func NewExpressionStatement(token *token.Token) *ExpressionStatement {
	return &ExpressionStatement{
		Token: token,
	}
}

func (s ExpressionStatement) statementNode() {}

func (s ExpressionStatement) TokenLiteral() string {
	return s.Token.Literal
}

func (s ExpressionStatement) String() string {
	if s.Expression != nil {
		return s.Expression.String()
	}

	return ""
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
