package lexer

import "github.com/sleepypioneer/monkey_interpreter/token"

// Transform source code -> tokens

// traverse source code calling NextToken() outputing token

// simplify by using string for source code type and not adding information such as filename or line number (as we would in production)

// lexer only supports ASCII characters instead of the full Unicode range.

type Lexer struct {
	input        string
	position     int  // position of current char
	readPosition int  // position after current char
	ch           byte // current char
}

//  New outputs new Lexer with input set from param
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // read first character and initialise position fields
	return l
}

//  advance our position in the input string
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for "NULL"
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
