package parser

import (
	"fmt"
	"pizzascript/lexer"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1 + 2", "{+,{1},{2}}"},
		{"1+2+3", "{+,{1},{+,{2},{3}}}"},
		{"1*2+3+4+5+1*2", "{+,{*,{1},{2}},{+,{3},{+,{4},{+,{5},{*,{1},{2}}}}}}"},
		{"1 + 2 * 3", "{+,{1},{*,{2},{3}}}"},
		{"1 * 2 + 3", "{+,{*,{1},{2}},{3}}"},
		{"1+2+3+4", "{+,{1},{+,{2},{+,{3},{4}}}}"},
		{"1*2+3*4", "{+,{*,{1},{2}},{*,{3},{4}}}"},
		{"1*2+3*4+5*6+7*8+9", "{+,{+,{+,{+,{*,{1},{2}},{*,{3},{4}}},{*,{5},{6}}},{*,{7},{8}}},{9}}"},
		{"1*2+3*4+5*6+7*8+9+100", "{+,{+,{+,{+,{*,{1},{2}},{*,{3},{4}}},{*,{5},{6}}},{*,{7},{8}}},{+,{9},{100}}}"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		fmt.Println(tt.input)
		actual := p.Print()

		if actual != tt.expected {
			t.Errorf("test failed for input %s, expected %s, actual %s", tt.input, tt.expected, actual)
		}
	}
}

func TestNudOperators(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"+1", "{1,{+}}"},
		{"++1", "{1,{+,{+}}}"},
		{"+++++1", "{1,{+,{+,{+,{+,{+}}}}}}"},
		{"-1", "{1,{-}}"},
		{"--1", "{1,{-,{-}}}"},
		{"+1+1", "{+,{1,{+}},{1}}"},
		{"+1++1", "{+,{1,{+}},{1,{+}}}"},
		{"+1++++++1", "{+,{1,{+}},{1,{+,{+,{+,{+,{+}}}}}}}"},
		{"+1+-+-+-1", "{+,{1,{+}},{1,{-,{+,{-,{+,{-}}}}}}}"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		fmt.Println(tt.input)
		actual := p.Print()

		if actual != tt.expected {
			t.Errorf("test failed for input %s, expected %s, actual %s", tt.input, tt.expected, actual)
		}
	}
}
