package lexer_test

import (
	"monkey/lexer"
	"monkey/token"
	"testing"
)

func TestLexerNextToken(t *testing.T) {
	input := `=+(){},;let`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.EOF, ""},
	}

	l := lexer.NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()

		//if err != nil {
		//	t.Fatalf("tests[%d] - error: %+v", i, err)
		//}

		if tok.Type == token.ILLEGAL {
			t.Fatalf("tests[%d] - illegal token: '%s'", i, tok.Literal)
		}

		if tok.Type != tt.expectedType {
			t.Errorf("tests[%d] - TokenType wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Errorf("tests[%d] - Literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
