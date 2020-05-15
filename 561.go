package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 1, 2, 3}
	a := arrayPairSum(nums)
	fmt.Println("a=", a)
}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	ret := 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			ret += nums[i]
		}
	}
	return ret
}
