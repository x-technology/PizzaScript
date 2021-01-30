package ast

import (
	"fmt"
	"pizzascript/token"
)

type Node struct {
	Token token.Token
	Left  *Node
	Right *Node
}

func (n *Node) ToString() {
	fmt.Print("{", n.Token.Literal)

	if n.Left != nil {
		fmt.Print(",")
		n.Left.ToString()
	}

	if n.Right != nil {
		fmt.Print(",")
		n.Right.ToString()
	}
	fmt.Print("}")
}

type Tree struct {
	root Node
}
