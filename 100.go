package main

import (
	"fmt"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {
	l1 := TreeNode{Val: 3}
	l1.Left = CreateNode(2)
	l2 := TreeNode{Val: 3}
	l2.Left = CreateNode(2)
	fmt.Println(isSameTree(&l1, &l2))
}

func CreateNode(v int) *TreeNode {
	return &TreeNode{Val: v}
}

func (node *TreeNode) SetValue(v int) {
	if node == nil {
		fmt.Println("setting value to nil.node ignored.")
		return
	}
	node.Val = v
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
