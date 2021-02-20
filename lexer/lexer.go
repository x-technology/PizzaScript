package lexer

import (
	"context"
	"fmt"
	"strings"
	"unicode"

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
			return !isWhitespace(str)
		}).
		Scan(func(_ context.Context, acc interface{}, next interface{}) (interface{}, error) {
			var tok interToken
			nextStr := next.(string)
			tok, isToken := acc.(interToken)
			tok.Return = ""

			// TODO change to types on string recognition
			if !isToken || (isNumber(tok.Save) && isNumber(nextStr)) {
				tok.Save += nextStr
			} else {
				tok.Return = tok.Save
				tok.Save = nextStr
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

func isWhitespace(ch string) bool {
	return ch == " " || ch == "\t" || ch == "\n" || ch == "\r"
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
	for isNumber(string(l.ch)) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || '0' <= ch && ch <= '9' || ch == '_'
}

// func isNumber(ch byte) bool {
// 	return '0' <= ch && ch <= '9'
// }

func isNumber(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
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

func isOperator(str string) bool {
	sort.Strings(token.ALL_OPERATORS)
	var index = sort.SearchStrings(token.ALL_OPERATORS, str)

	return token.ALL_OPERATORS[index] == str
}
