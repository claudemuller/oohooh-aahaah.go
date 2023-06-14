package lexer

import "github.com/claudemuller/oohooh-aahaah-go/internal/pkg/token"

/*
	Next Features:
		- Support Unicode
		- Support other number types
		- Add map, reduce etc. collection funcs
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
		tok = newToken(token.AssignTok, l.ch)

		if l.peekChar() == '=' {
			tok = token.Token{
				Type:    token.EQTok,
				Literal: l.getDoubleCharTok(),
			}
		}
	case '+':
		tok = newToken(token.PlusTok, l.ch)

	case '-':
		tok = newToken(token.MinusTok, l.ch)

	case '!':
		tok = newToken(token.BangTok, l.ch)

		if l.peekChar() == '=' {
			tok = token.Token{
				Type:    token.NotEQTok,
				Literal: l.getDoubleCharTok(),
			}
		}
	case '/':
		tok = newToken(token.SlashTok, l.ch)

	case '*':
		tok = newToken(token.AsteriskTok, l.ch)

	case '<':
		tok = newToken(token.LTTok, l.ch)

	case '>':
		tok = newToken(token.GTTok, l.ch)

	case ';':
		tok = newToken(token.SemicolonTok, l.ch)

	case ':':
		tok = newToken(token.ColonTok, l.ch)

	case ',':
		tok = newToken(token.CommaTok, l.ch)

	case '(':
		tok = newToken(token.LParenTok, l.ch)

	case ')':
		tok = newToken(token.RParenTok, l.ch)

	case '{':
		tok = newToken(token.LBraceTok, l.ch)

	case '}':
		tok = newToken(token.RBraceTok, l.ch)

	case '"':
		tok.Type = token.StringTok
		tok.Literal = l.readString()

	case '[':
		tok = newToken(token.LBracketTok, l.ch)

	case ']':
		tok = newToken(token.RBracketTok, l.ch)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOFTok

	default:
		tok = newToken(token.IllegalTok, l.ch)
		if isLetter(l.ch) {
			tok.Literal = l.readToken(isLetter) // Check if token is an identifier.
			tok.Type = token.LookupIdent(tok.Literal)

			return tok
		}

		if isDigit(l.ch) {
			tok.Type = token.IntTok
			tok.Literal = l.readToken(isDigit)

			return tok
		}
	}

	l.readChar()

	return tok
}

func (l *Lexer) readString() string {
	position := l.position + 1

	for {
		l.readChar()

		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[position:l.position]
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
	l.readPosition++
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

func (l *Lexer) getDoubleCharTok() string {
	ch := l.ch
	l.readChar()

	return string(ch) + string(l.ch)
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
