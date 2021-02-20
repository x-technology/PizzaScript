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
		" ":        true,
		"\n":       true,
		"\r":       true,
		"\t":       true,
		token.PLUS: false,
		"a":        false,
	}

	for i, isW := range tests {
		if isW != isWhitespace(i) {
			t.Fatalf("whitespace wrong. expected=%t, got=%t, where=%q", isW, isWhitespace(i), i)
		}
	}
}

func printTokens(input string) {
	l := New(input)
	l.Print()

	// for i, tt := range tests {
	// 	tok := l.NextToken()
	// 	if tok.Type != tt.expectedType {
	// 		t.Fatalf("[%d] - tokentype wrong. expected=%q, got=%q, where=%q",
	// 			i, tt.expectedType, tok.Type, tok.Literal)
	// 	}
	// 	if tok.Literal != tt.expectedLiteral {
	// 		t.Fatalf("[%d] - literal wrong. expected=%q, got=%q",
	// 			i, tt.expectedLiteral, tok.Literal)
	// 	}
	// 	t.Logf("[%d] ok - expected=%q, got=%q, where=%q",
	// 		i, tt.expectedType, tok.Type, tok.Literal)
	// }
}

func ExampleDifferentTokens(t *testing.T) {
	// t.Skip("not fully implemented")

	input := `=+(){},;512`
	printTokens(input)

	// Output
	// {= =}
	// {+ +}
	// {( (}
	// {) )}
	// {{ {}
	// {} }}
	// {, ,}
	// {; ;}
	// {INT 5}
}

func TestPizzaScript(t *testing.T) {
	t.Skip("not fully implemented")

	input := `var a1573: string = "null"`
	printTokens(input)
}

func TestVariableDeclarations(t *testing.T) {
	t.Skip("not fully implemented")

	input := `var h1679 = "1"
	val g2788 = 2
	`

	printTokens(input)
}

func TestIfElseStatements(t *testing.T) {
	t.Skip("not fully implemented")

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

	printTokens(input)
}
