package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 5, 6}
	target := 0
	a := searchInsert(nums, target)
	fmt.Println("a=", a)
}

func searchInsert(nums []int, target int) int {

	for i := 0; i < len(nums); i++ {
		if target < nums[i] && i == 0 {
			return 0
		} else if target == nums[i] {
			return i
		} else if target > nums[i] && i < len(nums)-1 && target < nums[i+1] {
			return i + 1
		} else if target > nums[i] && i == len(nums)-1 {
			return i + 1
		}
	}
	return 0
}
