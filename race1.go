package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 10}
	a := minCount(nums)
	fmt.Println("a=", a)
}

func minCount(coins []int) int {
	ret := 0

	for _, n := range coins {
		for n > 0 {
			n = n - 2
			ret++
		}
	}

	return ret
}
