package main

import (
	"fmt"
	"strconv"
)

func main() {
	ret := isPalindrome(1212)
	fmt.Println("ret-----> ", ret)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	} else {
		str := strconv.Itoa(x)

		for i := 0; i < len(str); i++ {
			if str[i] != str[len(str)-i-1] {
				return false
			}
		}
	}
	return true
}
