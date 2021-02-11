package lexer

import (
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

func (l *Lexer) NextToken() *token.Token {
	// 空白改行は読み飛ばす
	l.skipWhitespace()

	var tok *token.Token

	switch l.ch {
	case '=':
		tok = token.NewTokenByChar(token.ASSIGN, l.ch)
	case '+':
		tok = token.NewTokenByChar(token.PLUS, l.ch)
	case '-':
		tok = token.NewTokenByChar(token.MINUS, l.ch)
	case '!':
		tok = token.NewTokenByChar(token.BANG, l.ch)
	case '*':
		tok = token.NewTokenByChar(token.ASTERISK, l.ch)
	case '/':
		tok = token.NewTokenByChar(token.SLASH, l.ch)
	case '<':
		tok = token.NewTokenByChar(token.LT, l.ch)
	case '>':
		tok = token.NewTokenByChar(token.GT, l.ch)
	case '(':
		tok = token.NewTokenByChar(token.LPAREN, l.ch)
	case ')':
		tok = token.NewTokenByChar(token.RPAREN, l.ch)
	case '{':
		tok = token.NewTokenByChar(token.LBRACE, l.ch)
	case '}':
		tok = token.NewTokenByChar(token.RBRACE, l.ch)
	case ',':
		tok = token.NewTokenByChar(token.COMMA, l.ch)
	case ';':
		tok = token.NewTokenByChar(token.SEMICOLON, l.ch)
	case 0:
		tok = token.NewEOF()
	default:
		if l.isLetter() {
			// 識別子はreadIdentifierメソッド内で読み終わっているので、それ以上読む必要はない
			return l.readIdentifier()
		} else if l.isDigit() {
			// 数字はreadNumberメソッド内で読み終わっているので、それ以上読む必要はない
			return l.readNumber()
		}
		tok = token.NewTokenByChar(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return tok
}

// 識別子を読み進める
func (l *Lexer) readIdentifier() *token.Token {
	beginPosition := l.position
	for l.isLetter() {
		l.readChar()
	}
	literal := l.input[beginPosition:l.position]
	return token.NewIdentifierToken(literal)
}

// 使用可能な文字かチェックする
func (l *Lexer) isLetter() bool {
	return 'a' <= l.ch && l.ch <= 'z' || 'A' <= l.ch && l.ch <= 'Z' || l.ch == '_'
}

// 数字を読み進める
func (l *Lexer) readNumber() *token.Token {
	beginPosition := l.position
	for l.isDigit() {
		l.readChar()
	}
	literal := l.input[beginPosition:l.position]
	return token.NewIntegerToken(literal)
}

// 数字かチェックする
func (l *Lexer) isDigit() bool {
	return '0' <= l.ch && l.ch <= '9'
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

// 空白改行を無視する
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}
