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

func levelOrderBottom(root *Node) [][]int {
	// md := [][]int{}
	// if root == nil {
	// 	return md
	// }

	// layers := 0
	// nodes := []*Node{root}
	// mc := map[int][]int{}

	// mc[0] = append(mc[0], nodes[0].Value)
	// for len(nodes) > 0 {

	// 	layers++
	// 	mc[layers] = []int{}
	// 	size := len(nodes) //每层的节点数

	// 	count := 0
	// 	for count < size {

	// 		count++
	// 		curNode := nodes[0]
	// 		nodes = nodes[1:]
	// 		if curNode.Left != nil {
	// 			nodes = append(nodes, curNode.Left)
	// 			mc[layers] = append(mc[layers], curNode.Left.Value)
	// 		}
	// 		if curNode.Right != nil {
	// 			nodes = append(nodes, curNode.Right)
	// 			mc[layers] = append(mc[layers], curNode.Right.Value)
	// 		}
	// 	}
	// }

	// for i := len(mc) - 1; i >= 0; i-- {
	// 	v, ok := mc[i]
	// 	if ok && len(v) > 0 {
	// 		md = append(md, v)
	// 	}
	// }
	// return md

	result := [][]int{}
	if root == nil {
		return result
	}
	tmpResult := [][]int{}
	var q1, q2 []*Node
	q1 = append(q1, root)
	for len(q1) > 0 {
		q2 = q1
		q1 = []*Node{}
		for _, v := range q2 {
			if v.Left != nil {
				q1 = append(q1, v.Left)
			}
			if v.Right != nil {
				q1 = append(q1, v.Right)
			}
		}
		var tmp []int
		for _, v := range q2 {
			tmp = append(tmp, v.Value)
		}
		tmpResult = append(tmpResult, tmp)
	}
	for i := len(tmpResult) - 1; i >= 0; i-- {
		result = append(result, tmpResult[i])
	}
	return result
}

func CreateNode(v int) *Node {
	return &Node{Value: v}
}

func main() {
	root := Node{Value: 3}
	root.Left = CreateNode(9)
	root.Right = CreateNode(20)
	root.Right.Left = CreateNode(15)
	root.Right.Right = CreateNode(7)
	fmt.Println("\n层数: ", levelOrderBottom(&root))
}
