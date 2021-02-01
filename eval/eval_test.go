package eval

import (
	"fmt"
	"pizzascript/lexer"
	"pizzascript/parser"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"1 + 2", 3},
		{"1+2+3", 6},
		{"1*2+3+4+5+1*2", 16},
		{"1 + 2 * 3", 7},
		{"1 * 2 + 3", 5},
		{"1+2+3+4", 10},
		{"1*2+3*4", 14},
		{"1*2+3*4+5*6+7*8+9", 109},
		{"1*2+3*4+5*6+7*8+9+100", 209},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		e := New(p)

		fmt.Println(tt.input)
		actual := e.Eval()

		if actual != tt.expected {
			t.Errorf("test failed for input %s, expected %d, actual %d", tt.input, tt.expected, actual)
		}
	}
}
