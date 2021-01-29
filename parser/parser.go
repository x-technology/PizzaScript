package parser

import (
	"context"
	"pizzascript/ast"
	"pizzascript/lexer"
	"pizzascript/token"

	"github.com/reactivex/rxgo/v2"
)

type (
	prefixParseFn func(*ast.Node) *ast.Node
	infixParseFn  func(*ast.Node) *ast.Node
)

// Parser to parse PizzaScript tokens stream
// > A parser is a software component that takes input data (frequently text) and builds a data structure – often some kind of parse tree, abstract syntax tree or other hierarchical structure, giving a structural representation of the input while checking for correct syntax (c) Wikipedia
type Parser struct {
	lexer          *lexer.Lexer
	tree           rxgo.Observable
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

func nud(next token.Token) ast.Node {
	return ast.Node{Token: next}
}

func led(left ast.Node, operator token.Token, right ast.Node) ast.Node {
	return ast.Node{Left: &left, Token: operator, Right: &right}
}

type iterator struct {
	left     ast.Node
	operator *token.Token
}

func (p *Parser) Tree() rxgo.Observable {
	tree := p.lexer.Tokens().
		// 1 + 2     -> {1, +, 2}
		// 1 + 2 * 3 -> {1, +, {2, *, 3}}
		// 1 * 2 + 3 -> {{1, *, 2}, + , 3}
		Scan(func(_ context.Context, acc interface{}, elem interface{}) (interface{}, error) {
			it, isIterator := acc.(iterator)
			next := elem.(token.Token)

			// if acc is not defined, save nud and continue
			if !isIterator {
				it.left = nud(next)

				return it, nil
			}

			if it.operator != nil {
				// maybe not nud
				right := nud(next)
				it.left = led(it.left, *it.operator, right)
				it.operator = nil
			} else {
				it.operator = &next
			}

			return it, nil
		})

	return tree
}

func (p *Parser) Print() {
	for item := range p.tree.Observe() {
		it := item.V.(iterator)
		it.left.ToString()
		// fmt.Println(item.V)
	}
}
