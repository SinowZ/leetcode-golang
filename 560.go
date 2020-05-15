package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 4, 7, 2, -3, 1, 4, 2}
	num := subarraySum(nums, 7)
	fmt.Println("num=", num)
}

func subarraySum(nums []int, k int) int {
	sum, ans := 0, 0
	cnt := make(map[int]int)
	cnt[0] = 1
	for _, v := range nums {
		sum += v
		if tmp, ok := cnt[sum-k]; ok {
			ans += tmp

		}
		cnt[sum]++
	}

	return ans
}
