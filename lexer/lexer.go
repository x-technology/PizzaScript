package lexer

import (
	"context"
	"fmt"
	"strings"

	"pizzascript/token"
	"sort"

	"github.com/reactivex/rxgo/v2"
)

type Lexer struct {
	input        string
	observable   rxgo.Observable
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	observable := rxgo.Just(input)().FlatMap(func(i rxgo.Item) rxgo.Observable {
		splitted := strings.Split(i.V.(string), "")
		return rxgo.Just(splitted)()
	})

	l := &Lexer{input: input, observable: observable}
	return l
}

func (l *Lexer) Print() {
	for item := range l.Tokens().Observe() {
		fmt.Println(item.V)
	}
}

type interToken struct {
	Return string
	Save   string
}

func (l *Lexer) Tokens() rxgo.Observable {
	return rxgo.Concat([]rxgo.Observable{
		l.observable,
		// TODO fix when tokenize identifiers
		rxgo.Just("END")(),
	}).
		Filter(func(i interface{}) bool {
			var str = i.(string)
			ch := []byte(str)[0]
			return !isWhitespace(ch)
		}).
		Scan(func(_ context.Context, acc interface{}, elem interface{}) (interface{}, error) {
			var tok interToken
			tok, isToken := acc.(interToken)
			tok.Return = ""

			// TODO change to types on string recognition
			if !isToken || (isNumber([]byte(tok.Save)[0]) && isNumber([]byte(elem.(string))[0])) {
				tok.Save += elem.(string)
			} else {
				tok.Return = tok.Save
				tok.Save = elem.(string)
			}

			return tok, nil
		}).
		Filter(func(i interface{}) bool {
			tok := i.(interToken)

			return tok.Return != ""
		}).
		Map(func(_ context.Context, i interface{}) (interface{}, error) {
			var tok token.Token
			tok.Literal = i.(interToken).Return
			tok.Type = token.INT

			if isOperator(tok.Literal) {
				tok.Type = token.TokenType(tok.Literal)
			}

			return tok, nil
		})
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isNumber(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = "INT"
			return tok
		}
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func (l *Lexer) readString() string {
	l.readChar()
	position := l.position
	for l.ch != '"' {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isNumber(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || '0' <= ch && ch <= '9' || ch == '_'
}

func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isOperator(str string) bool {
	sort.Strings(token.ALL_OPERATORS)
	var index = sort.SearchStrings(token.ALL_OPERATORS, str)

	return token.ALL_OPERATORS[index] == str
}
