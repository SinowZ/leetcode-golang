package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
	a := findLHS(nums)
	fmt.Println("a=", a)
}

func findLHS(nums []int) int {
	count := 0
	m := map[int]int{}
	for _, n := range nums {
		m[n]++
	}

	for k, _ := range m {
		_, ok := m[k+1]
		if ok && m[k]+m[k+1] > count {
			count = m[k] + m[k+1]
		}
	}

	return count
}
