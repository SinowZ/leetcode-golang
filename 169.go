package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{4, 4, 4, 6, 6, 6, 9, 9}
	a := majorityElement(nums)
	fmt.Println("a=", a)
}

func majorityElement(nums []int) int {
	ret := math.MinInt64
	s := 0
	mc := map[int]int{}

	for _, n := range nums {
		if _, ok := mc[n]; !ok {
			mc[n] = 0
		} else {
			mc[n]++
		}
	}

	for k, v := range mc {
		if v > ret {
			ret = v
			s = k
		}
	}

	return s
}
