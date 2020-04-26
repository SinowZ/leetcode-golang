package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{12, 345, 2, 6, 7896}
	a := findNumbers(nums)
	fmt.Println("a=", a)
}

func findNumbers(nums []int) int {
	ret := 0

	for _, n := range nums {
		s := strconv.Itoa(n)
		if len(s)%2 == 0 {
			ret++
		}
	}

	return ret
}
