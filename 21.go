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
	l1.Insert(1, 1)
	l1.Insert(2, 2)
	l1.Insert(3, 4)
	l2 := &ListNode{0, nil}
	l2.Insert(1, 1)
	l2.Insert(2, 3)
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
	h := &ListNode{}
	l := h
	for nil != l1 && nil != l2 {
		if l1.Val <= l2.Val {
			l.Next = l1
			l1 = l1.Next
		} else {
			l.Next = l2
			l2 = l2.Next
		}
		l = l.Next
	}
	if l1 != nil {
		l.Next = l1
	} else {
		l.Next = l2
	}

	return h.Next
}
