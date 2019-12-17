package main

import (
	"fmt"
)

func main() {
	a := findTheDifference("abcd", "abcde")
	fmt.Println("a=", a)
}

func findTheDifference(s string, t string) byte {
	s = s + t
	var o byte

	for _, v := range []rune(s) {
		fmt.Println("v=", string(v))
		o ^= byte(v)
	}
	return o
}
