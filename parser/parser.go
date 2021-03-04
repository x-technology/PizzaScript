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

// TODO return iterator, differ in flags
func nud(acc interface{}, next token.Token) interface{} {
	// Node{Token: token.Token{Type: token.PLUS, Literal: "+"}, Left: &Node{Token: token.Token{Type: token.INT, Literal: "2"}}}
	node := ast.Node{Token: next}
	it := iterator{nud: &node}
	accIt, isIt := acc.(iterator)

	if isIt {
		// { , {... }
		node.Left = accIt.nud
	}

	// end
	if next.Type == token.INT {
		it.left = &node
		it.stack = accIt.stack
		it.nud = nil
	}
	
	return it
}

func led(left ast.Node, operator token.Token, right ast.Node) *ast.Node {
	return &ast.Node{Left: &left, Token: operator, Right: &right}
}

const (
	none = iota
	plus
	mul
)

func bp(tok *token.Token) int {
	if precendence, ok := precendences[string(tok.Type)]; ok {
		return precendence
	}

	return none
}

var precendences = map[string]int{
	token.PLUS:     plus,
	token.ASTERISK: mul,
}

type iterator struct {
	left     *ast.Node
	nud    	 *ast.Node
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
				// would stack grow in other dimensions?
				if len(it.stack) > 0 {
					prevIt = &it.stack[len(it.stack)-1]
				}

				for len(it.stack) > 0 && prevIt != nil && bp(it.operator) < bp(prevIt.operator) {
					prevIt = &it.stack[len(it.stack)-1]

					it.left = led(*prevIt.left, *prevIt.operator, *it.left)
					it.stack = it.stack[:len(it.stack)-1]
				}

				var newIt iterator
				// newIt.left = nud(nil, next).(iterator).left
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
