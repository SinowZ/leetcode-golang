package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 1, 2, 4}
	a := runningSum(nums)
	fmt.Println("a=", a)
}

func runningSum(nums []int) []int {
	mc := []int{}
	temp := 0
	for _, v := range nums {
		temp += v
		mc = append(mc, temp)
	}
	return mc
}
