package lexer

import "falcon/token"

type Lexer struct {
	input        string
	position     int  //current position in input (points to current char)
	readPosition int  //current reading position in input (after current char)
	ch           byte //current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		//ASCII code for "NUL"
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

//NextToken is the main function of the lexer. It returns the next token in the input string.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	//skip whitespace
	l.skipWhitespace()
	//switch statement to determine the type of the token
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
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
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	//ASCII code for "NUL"
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			//readIdentifier is a helper function that reads an identifier and advances the position of the lexer.
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			//readNumber is a helper function that reads a number and advances the position of the lexer.
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			//readNumber is a helper function that reads a number and advances the position of the lexer.
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}	
	l.readChar()
	return tok
}

//peekChar is a helper function that returns the next character in the input string without advancing the position of the lexer.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		//ASCII code for "NUL"
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

//readIdentifier is a helper function that reads an identifier and advances the position of the lexer.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

//isLetter is a helper function that checks if a character is a letter.
func isLetter(ch byte) bool {
	//ASCII code for "A" and "Z"
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

//readNumber is a helper function that reads a number and advances the position of the lexer.
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

//isDigit is a helper function that checks if a character is a digit.
func isDigit(ch byte) bool {
	//ASCII code for "0" and "9"
	return '0' <= ch && ch <= '9'
}

//newToken is a helper function that creates a new token.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

//skipWhitespace is a helper function that skips whitespace.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
