package parser

import (
	"context"
	"fmt"
	"pizzascript/ast"
	"pizzascript/lexer"
	"pizzascript/token"
)

type (
	prefixParseFn func(*ast.Node) *ast.Node
	infixParseFn  func(*ast.Node) *ast.Node
)

// Parser to parse PizzaScript tokens stream
// > A parser is a software component that takes input data (frequently text) and builds a data structure – often some kind of parse tree, abstract syntax tree or other hierarchical structure, giving a structural representation of the input while checking for correct syntax (c) Wikipedia
type Parser struct {
	lexer          *lexer.Lexer
	tree           *ast.Node
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{lexer: lexer}

	parser.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	parser.registerPrefix(token.INT, parser.parseIntegerLiteral)

	parser.tree = parser.Tree()

	return parser
}

func (p *Parser) parseIntegerLiteral(current *ast.Node) *ast.Node {
	return current
}

func nud(acc interface{}, next token.Token) iterator {
	// Node{Token: token.Token{Type: token.PLUS, Literal: "+"}, Left: &Node{Token: token.Token{Type: token.INT, Literal: "2"}}}
	node := ast.Node{Token: &next}
	it := iterator{nud: &node}
	accIt, isIt := acc.(iterator)

	if isIt {
		// { , {... }
		node.Left = accIt.nud
		it.stack = accIt.stack
	}

	// end
	if next.Type == token.INT {
		it.left = &node
		it.nud = nil
	}

	return it
}

func led(left ast.Node, operator token.Token, right ast.Node) *ast.Node {
	return &ast.Node{Left: &left, Token: &operator, Right: &right}
}

// iota enumerator gives an effective way in go language
// to use constants in enum-like constructs
const (
	none = iota
	plus
	mul
)

// calculate binding power
func bp(tok *token.Token) int {
	if precendence, ok := precendences[string(tok.Type)]; ok {
		return precendence
	}

	return none
}

// enum-like map of operator token keys and iota values defined above
var precendences = map[string]int{
	token.PLUS:     plus,
	token.ASTERISK: mul,
}

type iterator struct {
	left     *ast.Node
	nud      *ast.Node
	operator *token.Token
	stack    []iterator
}

func (p *Parser) Tree() *ast.Node {
	tree := p.lexer.Tokens().
		// 1 + 2     -> {1, +, 2}
		// 1 * 2 + 3 -> {{1, *, 2}, + , 3}
		// 1 + 2 * 3 -> {1, +, {2, *, 3}}
		// TODO change to reduce
		// TODO move to separate func
		// TODO make a stack interface
		Scan(func(_ context.Context, acc interface{}, elem interface{}) (interface{}, error) {
			it, isIterator := acc.(iterator)
			next := elem.(token.Token)

			// if acc is not defined, save nud and continue
			if !isIterator || it.left == nil {
				return nud(acc, next), nil
			}

			if it.operator != nil {
				var prevIt *iterator
				// TODO check if stack will grow in other dimensions and if that's a bug
				// try multiple binding power operators
				if len(it.stack) > 0 {
					prevIt = &it.stack[len(it.stack)-1]
				}

				for len(it.stack) > 0 && prevIt != nil && bp(it.operator) < bp(prevIt.operator) {
					prevIt = &it.stack[len(it.stack)-1]

					it.left = led(*prevIt.left, *prevIt.operator, *it.left)
					it.stack = it.stack[:len(it.stack)-1]
				}

				var newIt iterator
				newIt.stack = append(it.stack, it)

				return nud(newIt, next), nil
			}

			it.operator = &next
			return it, nil
		})

	lastItem, _ := tree.Last().Get()
	it := lastItem.V.(iterator)
	var prevIt *iterator

	for len(it.stack) > 0 {
		prevIt = &it.stack[len(it.stack)-1]

		it.left = led(*prevIt.left, *prevIt.operator, *it.left)
		it.stack = it.stack[:len(it.stack)-1]
	}

	return it.left
}

func (p *Parser) Print() string {
	lastItem := p.Tree()

	str := lastItem.ToString()
	fmt.Println(str)
	return str
}

func (p *Parser) PrintWat() string {
	lastItem := p.Tree()

	str := "(module\n  (func $add (result i32)\n    " + lastItem.ToWat() + "\n  )\n  (export \"add\" (func $add))\n)"
	fmt.Println(str)
	return str
}
