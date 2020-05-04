package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 1, 1, 4}
	a := jump(nums)
	fmt.Println("a=", a)
}

func jump(nums []int) int {
	position := len(nums) - 1
	steps := 0

	for position != 0 {
		for i := 0; i < position; i++ {
			if nums[i] >= position-i {
				position = i
				steps++
				break
			}
		}
	}
	return steps
}
