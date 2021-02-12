package token

import "fmt"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	detail  *DetailToken
}

func NewToken(tokenType TokenType, literal string) *Token {
	return &Token{
		Type:    tokenType,
		Literal: literal,
	}
}

func NewTokenByChar(tokenType TokenType, ch byte) *Token {
	return NewToken(tokenType, string(ch))
}

func NewIdentifierToken(literal string) *Token {
	tok := NewToken(lookupIdentifier(literal), literal)
	return tok
}

func NewIntegerToken(literal string) *Token {
	tok := NewToken(INT, literal)
	return tok
}

func NewEOF() *Token {
	return &Token{Type: EOF, Literal: ""}
}

func (t *Token) SetDetail(detail *DetailToken) {
	t.detail = detail
}

func (t *Token) IsEOF() bool {
	return t.Type == EOF
}

func (t *Token) Detail() *DetailToken {
	return t.detail
}

func (t *Token) Debug() string {
	if t.Type == IDENT || t.Type == INT {
		return fmt.Sprintf("%s(%q)", t.Type, t.Literal)
	}
	return fmt.Sprintf("%q", t.Literal)
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 識別子 + リテラル
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// 演算子
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func lookupIdentifier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENT
}
