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

type LetStatement struct {
	Token *token.Token // token.LET トークン
	Name  *Identifier
	Value Expression
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

type ReturnStatement struct {
	Token       *token.Token // token.RETURN トークン
	ReturnValue Expression
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

type Identifier struct {
	Token *token.Token // token.IDENT トークン
	Value Expression
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
