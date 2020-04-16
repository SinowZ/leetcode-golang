package main

import (
	"fmt"
)

func main() {
	a := isPowerOfThree(1)
	fmt.Println("a=", a)
}

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	if n < 3 {
		return false
	}

	for n >= 3 {
		if n%3 != 0 {
			return false
		}
		n = n / 3
		if n < 3 {
			if n == 1 {
				return true
			} else {
				return false
			}
		}
	}

	return false
}
