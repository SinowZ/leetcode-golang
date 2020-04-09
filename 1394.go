package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{4, 6, 6, 6, 2, 2}
	a := findLucky(arr)
	fmt.Println("a=", a)
}

func findLucky(arr []int) int {
	ret := []int{}
	mc := map[int]int{}

	for n := range arr {
		_, ok := mc[arr[n]]
		if !ok {
			mc[arr[n]] = 1
		} else {
			mc[arr[n]]++
		}
	}

	for k, v := range mc {
		if k == v {
			ret = append(ret, k)
		}
	}

	if len(ret) == 0 {
		return -1
	}

	max := math.MinInt64
	for n := range ret {
		if ret[n] > max {
			max = ret[n]
		}
	}
	return max
}
