package main

import (
	"fmt"
)

func main() {
	ret := divide(10, 3)
	fmt.Println("ret= ", ret)
}

func divide(dividend int, divisor int) int {
	if dividend == 0 || divisor == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 && dividend == -1<<31 {
		return 1<<31 - 1
	} else if divisor == -1 {
		return 0 - dividend
	}

	isFu := false
	if dividend < 0 && divisor > 0 {
		isFu = true
		dividend = 0 - dividend
	}
	if dividend > 0 && divisor < 0 {
		isFu = true
		divisor = 0 - divisor
	}
	if dividend < 0 && divisor < 0 {
		divisor = 0 - divisor
		dividend = 0 - dividend
	}

	temp := dividend
	ret := 0
	for temp >= divisor {
		temp -= divisor
		ret++
	}

	if isFu {
		return 0 - ret
	}
	return ret
}
