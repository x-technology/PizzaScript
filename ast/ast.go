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

type Tree struct {
	root Node
}
