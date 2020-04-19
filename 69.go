package main

import (
	"fmt"
)

func main() {
	a := mySqrt(4)
	fmt.Println("a=", a)
}

func mySqrt(x int) int {
	ret := 0

	if x == 0 {
		ret = 0
	} else if x == 1 || x == 2 {
		ret = 1
	} else {
		for i := 1; i <= x/2; i++ {
			if i*i <= x && (i+1)*(i+1) > x {
				return i
			}
		}
	}
	return ret
}
