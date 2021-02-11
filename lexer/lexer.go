package lexer

import (
	"errors"
	"fmt"
	"monkey/token"
)

type Lexer struct {
	input        string // 字句解析対象の入力文字列
	position     int    // 入力における現在の位置／現在の文字を指し示す
	readPosition int    // これから読み込む位置／現在の文字の次
	ch           byte   // 現在検査中の文字
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() (*token.Token, error) {
	var tok *token.Token

	switch l.ch {
	case '=':
		tok = token.NewToken(token.ASSIGN, l.ch)
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)
	case 0:
		tok = token.NewEOF()
	default:
		message := fmt.Sprintf("error NextToken: undefined token: '%s'\n", string(l.ch))
		return nil, errors.New(message)
	}

	l.readChar()
	return tok, nil
}

// 次の一文字を読んで、位置ポインタを更新する
// positionは常にreadPositionの次を指し示す
// 終端までいったらASCIIコードのNUL文字をセットする
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // NUL文字
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
