package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 4, 11, 3, 3, 6, 6, 9, 9}
	a := singleNumber(nums)
	fmt.Println("a=", a)
}

func singleNumber(nums []int) int {
	o := 0
	for _, n := range nums {
		o ^= n
		fmt.Println("o=", o)
	}
	return o
}
