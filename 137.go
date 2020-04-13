package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 2, 3, 2}
	a := findMin(nums)
	fmt.Println("a=", a)
}

//l^n 不进位加法计算  4^2 = 0000 0100 + 0000 0010 = 0000 0110 = 6
//^h  按位取反       ^4 = 0000 0100
// & 相加
func findMin(nums []int) int {
	h, l := 0, 0
	for _, n := range nums {
		l = (l ^ n) & (^h)
		h = (h ^ n) & (^l)
	}
	return l
}
