package main

import "fmt"

func main() {
	val := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4}
	length := removeDuplicates(val)
	fmt.Println("length-----", length)
}

func removeDuplicates(nums []int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != nums[i] {
			nums[i+1] = nums[j]
			i++
		}
	}

	return i + 1
}
