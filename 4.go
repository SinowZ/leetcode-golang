package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 4}
	a := findMedianSortedArrays(nums1, nums2)
	fmt.Println("a=", a)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	for v := range nums2 {
		nums1 = append(nums1, nums2[v])
	}

	for i := 0; i < len(nums1)-1; i++ {
		for j := i + 1; j < len(nums1); j++ {
			if nums1[i] > nums1[j] {
				nums1[i], nums1[j] = nums1[j], nums1[i]
			}
		}
	}

	if len(nums1)%2 == 0 {
		leftIndex := len(nums1)/2 - 1
		rightIndex := leftIndex + 1
		return float64(nums1[leftIndex]+nums1[rightIndex]) / 2
	}
	if len(nums1)%2 == 1 {
		middlendex := len(nums1) / 2
		return float64(nums1[middlendex])
	}

	return 0
}
