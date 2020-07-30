package main

import (
	"fmt"
)

func main() {
	a := findLUSlength("aba", "cdc")
	fmt.Println("a=", a)
}

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	return max(len(a), len(b))
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
