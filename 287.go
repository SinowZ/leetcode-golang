package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 4, 6, 6, 8, 9}
	a := findDuplicate(nums)
	fmt.Println("a=", a)
}

func findDuplicate(nums []int) int {
	mc := map[int]int{}

	for _, n := range nums {
		if _, ok := mc[n]; !ok {
			mc[n] = 0
		} else {
			return n
		}
	}
	return nums[0]
}
