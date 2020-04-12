package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 4, 6, 6, 9, 9}
	a := findMin(nums)
	fmt.Println("a=", a)
}

func findMin(nums []int) int {
	ret := nums[0]
	for n := range nums {
		if nums[n] < ret {
			ret = nums[n]
		}
	}
	return ret
}
