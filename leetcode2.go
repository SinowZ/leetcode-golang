package main

import (
	"fmt"
	"math"
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

/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret = &ListNode{0, nil}
	var p = ret
	var carry = 0
	for l1 != nil || l2 != nil || carry != 0 {
		var intL1, a2 int
		if l1 != nil {
			intL1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			a2 = l2.Val
			l2 = l2.Next
		}
		sum := intL1 + a2 + carry
		carry = sum / 10
		fmt.Println(sum % 10)
		p.Next = &ListNode{sum % 10, nil}
		p = p.Next
	}
	return ret.Next
}
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	ret := &ListNode{0, nil}
	p := ret

	var intL1, intL2 int

	point1 := l1.Next
	i := 0
	for nil != point1 {

		if intL1 == 0 {
			intL1 = point1.Val
		} else {
			intL1 = point1.Val*int(math.Pow10(i)) + intL1
		}
		i++
		point1 = point1.Next
	}

	i = 0
	point2 := l2.Next
	for nil != point2 {
		if intL2 == 0 {
			intL2 = point2.Val
		} else {
			intL2 = point2.Val*int(math.Pow10(i)) + intL2
		}
		i++
		point2 = point2.Next
	}
	i = intL1 + intL2

	//p.Next = &ListNode{0, nil}
	//p = p.Next

	for i > 0 {
		p.Next = &ListNode{i % 10, nil}
		p = p.Next
		i = i / 10
	}

	return ret.Next
}
