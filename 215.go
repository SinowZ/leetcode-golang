package main

import (
	"fmt"
	"sort"
)

func main() {
	a := findKthLargest([]int{1, 2, 3, 1}, 3)
	fmt.Println("a=", a)
}

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
