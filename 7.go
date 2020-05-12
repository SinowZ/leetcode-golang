package main

import (
	"fmt"
)

func main() {
	ret := reverse(12345)
	fmt.Println("ret=", ret)
}

func reverse(x int) int {
	result := 0

	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}

	if result > 2147483647 || result < -2147483648 {
		return 0
	}
	return result
}
