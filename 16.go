package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	target := 6
	ret := threeSumClosest(nums, target)
	fmt.Println("ret= ", ret)
}

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	sort.Ints(nums)
	diff := 0
	for i := 0; i < n; i++ {
		if i == 0 {
			diff = nums[0] + nums[1] + nums[n-1] - target
		} else if nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			temp := nums[i] + nums[j] + nums[k] - target
			if temp == 0 {
				return target
			} else if temp < 0 {
				j++
			} else {
				k--
			}
			if abs(temp) < abs(diff) {
				diff = temp
			}
		}
	}
	return diff + target
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
