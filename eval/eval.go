package eval

import (
	"pizzascript/ast"
	"pizzascript/parser"
	"strconv"

	"pizzascript/utils/log"
)

type Evaluator struct {
	parser *parser.Parser
}

func New(parser *parser.Parser) *Evaluator {
	eval := &Evaluator{parser: parser}

	return eval
}

func (e *Evaluator) Eval() int64 {
	lastItem := e.parser.Tree()
	value := eval(lastItem)

	return value
}

func isNud(node *ast.Node) bool {
	return node.Left != nil && node.Token != nil && node.Right == nil
}

func eval(node *ast.Node) int64 {
	value := node.Token.Literal

	if isNud(node) {
		right := evalIntegerLiteral(value)
		
		for node != nil && node.Left != nil {
			operator := node.Left.Token.Literal
			right = evalIntegerInfixExpression(operator, 0, right)

			node = node.Left
		}

		return right
	}

	if node.Left == nil && node.Right == nil {
		return evalIntegerLiteral(value)
	}

	left := eval(node.Left)
	right := eval(node.Right)

	return evalIntegerInfixExpression(value, left, right)
}

func evalIntegerPrefixExpression(
	operator string,
	left int64,
) int64 {
	switch operator {
	case "+":
		return left
	case "-":
		return -left
	default:
		log.Err("unknown operator", left, operator)
		return 0
	}
}

func evalIntegerInfixExpression(
	operator string,
	left, right int64,
) int64 {
	switch operator {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	default:
		log.Err("unknown operator", left, operator, right)
		return 0
	}
}

func evalIntegerLiteral(str string) int64 {
	value, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		log.Err("could not parse as integer", str)
	}

	return value
}
