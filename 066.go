package main

import (
	"fmt"
)

func main() {
	digits := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	a := plusOne(digits)
	fmt.Println("a=", a)
}

func plusOne(digits []int) []int {
	acc := 1
	for i := len(digits) - 1; i >= 0 && acc == 1; i-- {
		n := digits[i] + 1
		if n == 10 {
			n = 0
		} else {
			acc = 0
		}
		digits[i] = n
	}
	if acc == 0 {
		return digits
	}
	digits = append(digits, 0)
	copy(digits[1:], digits)
	digits[0] = 1
	return digits
}
