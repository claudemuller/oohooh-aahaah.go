package lexer

import "github.com/claudemuller/oohooh-aahaah-go/internal/pkg/token"

/*
	Next Features:
		- Support Unicode
		- Support other number types
*/

type Lexer struct {
	input        string
	position     int  // the current position in the "input" (current char).
	readPosition int  // the current reading position in "input" (after current char).
	ch           byte // the current char under examination.
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.eatWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)

		// TODO: Improve this duplication.
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{
				Type:    token.EQ,
				Literal: literal,
			}
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		tok = newToken(token.BANG, l.ch)

		// TODO: Improve this duplication.
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{
				Type:    token.NOT_EQ,
				Literal: literal,
			}
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		tok = newToken(token.ILLEGAL, l.ch)
		if isLetter(l.ch) {
			tok.Literal = l.readToken(isLetter) // Check if token is an identifier.
			tok.Type = token.LookupIdent(tok.Literal)

			return tok
		}

		if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readToken(isDigit)

			return tok
		}
	}

	l.readChar()

	return tok
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readToken(isFn func(byte) bool) string {
	position := l.position
	for isFn(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readChar() {
	l.ch = 0 // ASCII char for NUL.
	if l.readPosition < len(l.input) {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition < len(l.input) {
		return l.input[l.readPosition]
	}
	return 0
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
