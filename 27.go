package main

import (
	"fmt"
)

func main() {
	//nums := []int{1, 2, 2, 3, 2, 3, 4, 5, 6, 2}
	nums := []int{1}
	a := removeElement(nums, 1)
	fmt.Println("a=", a)
}

func removeElement(nums []int, val int) int {
	i, j := 0, 0
	for ; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		nums[j] = nums[i]
		j++
	}
	return j
}
