package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 6, 6, 6, 9, 9}
	nums2 := []int{6, 6, 4, 9}
	a := intersect(nums, nums2)
	fmt.Println("a=", a)
}

func intersect(nums1 []int, nums2 []int) []int {
	ret := []int{}
	mc := map[int]int{}

	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	for _, n := range nums1 {
		_, ok := mc[n]
		if !ok {
			mc[n] = 0
		} else {
			mc[n]++
		}
	}

	for _, n := range nums2 {
		if _, ok := mc[n]; ok {
			ret = append(ret, n)
			if mc[n] > 0 {
				mc[n]--
			} else {
				delete(mc, n)
			}
		}
	}
	return ret
}
