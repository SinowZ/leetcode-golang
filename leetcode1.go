package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 2, 4, -2, -1, 0, -3}
	//target := 0
	//fmt.Println(twoSum(nums, target))
	fmt.Println(threeSum(nums))
}

func twoSum(nums []int, target int) []int {
	result := []int{}

	for i, valueLeft := range nums {
		for j, valueRight := range nums {
			if target-valueLeft == valueRight && i != j {
				result = append(result, i)
				result = append(result, j)
				return result
			}
		}
	}
	return result
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums)-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		z := len(nums) - 1
		for z > j {
			b := nums[j]
			c := nums[z]
			if nums[i]+b+c > 0 {
				z--
			} else if nums[i]+b+c < 0 {
				j++
			} else {
				item := []int{nums[i], b, c}
				result = append(result, item)
				for j < z && nums[j] == nums[j+1] {
					j++
				}
				for j < z && nums[z] == nums[z-1] {
					z--
				}
				j++
				z--
			}
		}
	}
	return result
}
