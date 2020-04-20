package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	temp := 0
	for i, _ := range nums1 {
		if i >= m {
			nums1[i] = nums2[temp]
			temp++
		}
	}

	for i := 0; i < len(nums1)-1; i++ {
		for j := i + 1; j < len(nums1); j++ {
			if nums1[i] > nums1[j] {
				nums1[i], nums1[j] = nums1[j], nums1[i]
			}
		}
	}
	fmt.Println(nums1)
}
