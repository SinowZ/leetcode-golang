package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := addDigits(38)
	fmt.Println("a=", a)
}

func addDigits(num int) int {
	if num < 10 {
		return num
	}

	for num >= 10 {
		temp := 0
		for _, s := range strconv.Itoa(num) {
			n, _ := strconv.Atoi(string(s))
			temp += n
		}
		num = temp
	}
	return num
}
