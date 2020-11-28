package lexer

import (
	"github.com/pizzascript/pizzascript/token"

	"testing"
)

type lexerTests []struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func runTests(t *testing.T, input string, tests lexerTests) {
	l := New(input)
	t.Logf("Test [%q]:", input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("[%d] - tokentype wrong. expected=%q, got=%q, where=%q",
			i, tt.expectedType, tok.Type, tok.Literal)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("[%d] - literal wrong. expected=%q, got=%q",
			i, tt.expectedLiteral, tok.Literal)
		}
		t.Logf("[%d] ok - expected=%q, got=%q, where=%q",
		i, tt.expectedType, tok.Type, tok.Literal)
	}
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := lexerTests{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	runTests(t, input, tests)
}

func TestPizzaScript(t *testing.T) {
	input := `var a1573: string = "null"`
	tests := lexerTests{
		{token.LET, "var"},
		{token.IDENT, "a1573"},
		{token.COLON, ":"},
		{token.TYPE, "string"},
		{token.ASSIGN, "="},
		{token.STRING, "null"},
		{token.EOF, ""},
	}

	runTests(t, input, tests)
}

func TestVariableDeclarations(t *testing.T) {
	input := `var h1679 = "1"
	val g2788 = 2
	`
	tests := lexerTests{
		{token.LET, "var"},
		{token.IDENT, "h1679"},
		{token.ASSIGN, "="},
		{token.STRING, "1"},
		{token.LET, "val"},
		{token.IDENT, "g2788"},
		{token.ASSIGN, "="},
		{token.INT, "2"},
		{token.EOF, ""},
	}

	runTests(t, input, tests)
}
