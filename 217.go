package main

import (
	"fmt"
)

func main() {
	a := containsDuplicate([]int{1, 3})
	fmt.Println("a=", a)
}

func containsDuplicate(nums []int) bool {
	m := map[int]int{}
	for _, n := range nums {
		if m[n] > 0 {
			return true
		}
		m[n]++
	}
	return false
}
