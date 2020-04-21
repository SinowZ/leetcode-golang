package main

import (
	"fmt"
)

func main() {
	a := sortedSquares([]int{-4, -1, 0, 3, 10})
	fmt.Println("a=", a)
}

func sortedSquares(A []int) []int {
	for i, v := range A {
		A[i] = v * v
		if A[i] < 0 {
			A[i] = 0 - A[i]
		}
	}

	for i := 0; i < len(A)-1; i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i] > A[j] {
				A[i], A[j] = A[j], A[i]
			}
		}
	}
	return A
}
