package ast

import (
	"fmt"
	"pizzascript/token"
	"testing"
)

func TestToString(t *testing.T) {
	tests := []struct {
		input    Node
		expected string
	}{
		{Node{Token: token.Token{Type: token.INT, Literal: "1"}}, "{1}"},
		{Node{Token: token.Token{Type: token.PLUS, Literal: "+"}, Left: &Node{Token: token.Token{Type: token.INT, Literal: "2"}}}, "{+,{2}}"},
	}

	for _, tt := range tests {
		fmt.Println(tt.input)
		actual := tt.input.ToString()

		if actual != tt.expected {
			t.Errorf("test failed for expected %s, actual %s", tt.expected, actual)
		}
	}
}