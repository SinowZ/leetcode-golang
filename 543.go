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

var steps int

func diameterOfBinaryTree(root *Node) int {
	steps = 1
	depth(root)
	return steps - 1
}

func depth(root *Node) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left)
	right := depth(root.Right)
	if left+right+1 > steps {
		steps = left + right + 1
	}
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

func main() {
	root := CreateNode(1)
	root.Left = CreateNode(2)
	root.Left.Left = CreateNode(4)
	root.Left.Right = CreateNode(5)
	root.Right = CreateNode(5)
	root.PreOrder()
	fmt.Println()
	s := diameterOfBinaryTree(root)
	fmt.Println(s)
}
