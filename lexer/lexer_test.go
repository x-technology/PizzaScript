package lexer

import (
	"pizzascript/token"

	"testing"
)

type lexerTests []struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestIsWhiteSpace(t *testing.T) {
	tests := map[string]bool{
		" " : true,
		"\n" : true,
		"\r" : true,
		"\t" : true,
		token.PLUS: false,
		"a": false,
	}

	for i, isW := range tests {
		if isW != isWhitespace(i) {
			t.Fatalf("whitespace wrong. expected=%t, got=%t, where=%q", isW, isStrWhitespace(i), i)
		}
	}
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

func TestIfElseStatements(t *testing.T) {
	input :=
		`if 5 < 10 {
				true
		} else {
				false
		}`

	tests := lexerTests{
		{token.IF, "if"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.LBRACE, "{"},
		{token.TRUE, "true"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.FALSE, "false"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	runTests(t, input, tests)
}
