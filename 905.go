package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 1, 2, 4}
	a := sortArrayByParity(nums)
	fmt.Println("a=", a)
}

func sortArrayByParity(A []int) []int {
	mc := []int{}
	md := []int{}
	for n := range A {
		if A[n]%2 == 0 {
			mc = append(mc, A[n])
			fmt.Println("mc------", A[n])
		}
		if A[n]%2 == 1 {
			md = append(md, A[n])
			fmt.Println("md------", A[n])
		}
	}
	return append(mc, md...)
}
