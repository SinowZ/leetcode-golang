package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{0, nil}
	l1.Insert(1, 2)
	l1.Insert(2, 4)
	l1.Insert(3, 3)
	l2 := &ListNode{0, nil}
	l2.Insert(1, 5)
	l2.Insert(2, 6)
	l2.Insert(3, 4)
	fmt.Println(addTwoNumbers(l1, l2))
}

func (head *ListNode) Insert(i int, e int) bool {
	p := head
	j := 1
	for nil != p && j < i {
		p = p.Next
		j++
	}
	if nil == p || j > i {
		fmt.Println("pls check i:", i)
		return false
	}
	s := &ListNode{Val: e}
	s.Next = p.Next
	p.Next = s
	return true
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret = &ListNode{0, nil}
	var p = ret
	var i = 0
	for l1 != nil || l2 != nil || i != 0 {
		var intL1, intL2 int
		if l1 != nil {
			intL1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			intL2 = l2.Val
			l2 = l2.Next
		}
		total := intL1 + intL2 + i
		i = total / 10
		fmt.Println(total % 10)
		p.Next = &ListNode{total % 10, nil}
		p = p.Next
	}
	return ret.Next
}
