package ast

import (
	"pizzascript/token"
)

type Node struct {
	Token *token.Token
	Left  *Node
	Right *Node
}

func (n *Node) ToString() string {
	res := "{" + n.Token.Literal

	if n.Left != nil {
		res += "," + n.Left.ToString()
	}

	if n.Right != nil {
		res += "," + n.Right.ToString()
	}

	res += "}"
	return res
}

//
// (module
//   (func $add (result i32)
//     i32.const 13
//     i32.const 13
//     i32.add
//     i32.const 13
//     i32.const 13
//     i32.add
//     i32.add)
//   (export "add" (func $add))
// )
func (n *Node) ToWat() string {
	res := ""

	if n.Left != nil {
		res += n.Left.ToWat()
	}

	if n.Right != nil {
		res += n.Right.ToWat()
	}

	if n.Token.Type == token.INT {
		res += "i32.const " + n.Token.Literal
	} else {
		switch n.Token.Literal {
		case "+":
			res += "i32.add"
		case "-":
			res += "i32.sub"
		case "*":
			res += "i32.mul"
		case "/":
			res += "i32.div_s"
		}
	}

	res += "\n"
	return res
}
