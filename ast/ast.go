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
	fmt.Println("token", n.Token)

	if n.Left != nil {
		fmt.Print("left")
		n.Left.ToString()
	}

	if n.Right != nil {
		fmt.Print("right")
		n.Right.ToString()
	}
}

type Tree struct {
	root Node
}
