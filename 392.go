package main

import (
	"fmt"
	"strings"
)

func main() {
	a := isSubsequence("abc", "ahbgdc")
	fmt.Println("a=", a)
}

func isSubsequence(s string, t string) bool {
	for _, v := range []rune(s) {
		location := strings.Index(t, string(v))
		if location == -1 {
			return false
		}
		t = t[location+1 : len(t)]
	}
	return true
}
