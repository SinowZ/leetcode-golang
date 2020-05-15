package main

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

func (node *Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(v int) {
	if node == nil {
		fmt.Println("setting value to nil.node ignored.")
		return
	}
	node.Value = v
}

func CreateNode(v int) *Node {
	return &Node{Value: v}
}

//前序遍历
func (node *Node) PreOrder() {
	if node == nil {
		return
	}
	node.Print()
	node.Left.PreOrder()
	node.Right.PreOrder()
}

func rangeSumBST(root *Node, L int, R int) int {
	if root == nil {
		return 0
	}
	if root.Value < L {
		return rangeSumBST(root.Right, L, R)
	}
	if root.Value > R {
		return rangeSumBST(root.Left, L, R)
	}
	return root.Value + rangeSumBST(root.Right, L, R) + rangeSumBST(root.Left, L, R)
}

func main() {
	root := CreateNode(10)
	root.Left = CreateNode(5)
	root.Left.Left = CreateNode(2)
	root.Left.Right = CreateNode(7)
	root.Right = CreateNode(15)
	root.Right.Right = CreateNode(18)
	root.PreOrder()
	fmt.Println()
	s := rangeSumBST(root, 7, 15)
	fmt.Println(s)
}
