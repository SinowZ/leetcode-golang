package main

import (
	"fmt"
)

func main() {
	nums := []int{99, 99}
	a := containsNearbyDuplicate(nums, 2)
	fmt.Println("a=", a)
}

func containsNearbyDuplicate(nums []int, k int) bool {

	lo, j, hi := 0, 1, len(nums)-1
	for lo <= hi {
		for j <= k && lo+j <= hi {
			//fmt.Println(lo, lo+j, nums[lo], nums[lo+j])
			if nums[lo] == nums[lo+j] {
				return true
			} else {
				j++
			}
		}
		lo++
		j = 1
	}
	return false
}
