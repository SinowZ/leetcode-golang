package main

import (
	"fmt"
	"math"
)

func main() {
	a := containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0)
	fmt.Println("a=", a)
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if int(math.Abs(float64(nums[i]-nums[j]))) <= t && int(math.Abs(float64(i-j))) <= k {
				return true
			}

		}
	}
	return false
}
