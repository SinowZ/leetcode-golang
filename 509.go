package main

import (
	"fmt"
)

func main() {
	num := fib(3)
	fmt.Println("num=", num)
}

func fib(N int) int {
	ret := 0
	if N == 0 {
		ret = 0
	} else if N == 1 {
		ret = 1
	} else {
		ret = fib(N-1) + fib(N-2)
	}
	return ret
}
