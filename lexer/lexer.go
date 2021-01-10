package lexer

import (
	"context"
	"fmt"
	"strings"
	// "sort"
	"pizzascript/token"
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

func (l *Lexer) Tokens() rxgo.Observable {
	return rxgo.Merge([]rxgo.Observable{
		// operators
		// l.observable.
		// Filter(func(i interface{}) bool {
		// 	var str = i.(string)

		// 	sort.Strings(token.ALL_OPERATORS)
		// 	var index = sort.SearchStrings(token.ALL_OPERATORS, str)

		// 	return token.ALL_OPERATORS[index] == str
		// }).
		// Map(func(_ context.Context, i interface{})(interface{}, error) {
		// 	var str = i.(string)
		// 	ch := []byte(str)[0]
		// 	var index = sort.SearchStrings(token.ALL_OPERATORS, str)

		// 	var tok token.Token = newToken(token.TokenType(token.ALL_OPERATORS[index]), ch)
		// 	return tok, nil
		// }),
		// numbers
		l.observable.
		// Filter(func(i interface{}) bool {
		// 	var str = i.(string)
		// 	ch := []byte(str)[0]

		// 	return isNumber(ch)
		// }).
		DistinctUntilChanged(func(_ context.Context, i interface{}) (interface{}, error) {
			var str = i.(string)
			ch := []byte(str)[0]
			fmt.Println(str)
	
			return isNumber(ch), nil
		}).
		// Reduce(func(_ context.Context, acc interface{}, elem interface{}) (interface{}, error) {
		// 	if acc == nil {
		// 		return elem, nil
		// 	}
		// 	return acc.(string) + elem.(string), nil
		// }),
		// SumInt64()
		Map(func(_ context.Context, i interface{})(interface{}, error) {
			var str = i.(string)
			ch := []byte(str)[0]

			var tok token.Token = newToken(token.INT, ch)
			return tok, nil
		}),
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
