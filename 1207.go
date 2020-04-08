package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 4, 6, 6, 9, 9}
	a := uniqueOccurrences(nums)
	fmt.Println("a=", a)
}

func uniqueOccurrences(arr []int) bool {
	mc := map[int]int{}
	md := map[int]int{}

	for _, n := range arr {
		_, ok := mc[n]
		if !ok {
			mc[n] = 0
		} else {
			mc[n]++
		}
	}

	for _, v := range mc {
		_, ok := md[v]
		if ok {
			return false
		} else {
			md[v] = 0
		}
	}

	return true
}
