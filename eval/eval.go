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
	value := eval(&lastItem)

	return value
}

func eval(node *ast.Node) int64 {
	value := node.Token.Literal
	if node.Left == nil && node.Right == nil {
		return evalIntegerLiteral(value)
	}

	left := eval(node.Left)
	right := eval(node.Right)

	return evalIntegerInfixExpression(value, left, right)
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
