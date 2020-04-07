package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 4, 6, 6, 6, 9, 9}
	nums2 := []int{6, 4, 5, 9}
	a := intersection(nums, nums2)
	fmt.Println("a=", a)
}

func intersection(nums1 []int, nums2 []int) []int {
	ret := []int{}
	mc := map[int]int{}
	md := map[int]int{}

	for _, n := range nums1 {
		if _, ok := mc[n]; !ok {
			mc[n] = 0
		}
	}

	for _, n := range nums2 {
		if _, ok := mc[n]; ok {
			md[n] = 0
		}
	}

	for k, _ := range md {
		ret = append(ret, k)
	}

	return ret
}
