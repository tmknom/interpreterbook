package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, literal string) *Token {
	return &Token{Type: tokenType, Literal: literal}
}

func NewTokenByChar(tokenType TokenType, ch byte) *Token {
	return NewToken(tokenType, string(ch))
}

func NewIdentifierToken(literal string) *Token {
	return &Token{Type: lookupIdentifier(literal), Literal: literal}
}

func NewIntegerToken(literal string) *Token {
	return &Token{Type: INT, Literal: literal}
}

func NewEOF() *Token {
	return &Token{Type: EOF, Literal: ""}
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 識別子 + リテラル
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// 演算子
	ASSIGN = "="
	PLUS   = "+"

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
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func lookupIdentifier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENT
}
