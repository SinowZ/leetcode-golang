package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	a := maxSubArray(nums)
	fmt.Println("a=", a)
}

func maxSubArray(nums []int) int {
	sum := nums[0]
	now := nums[0]

	for i := 1; i < len(nums); i++ {
		if now < 0 {
			now = nums[i]
		} else {
			now = now + nums[i]
		}
		if now > sum {
			sum = now
		}
	}
	return sum
}
