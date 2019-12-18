package main

import (
	"fmt"
)

func main() {
	s := "A man, a plan, a canal: Panamaa"
	a := isPalindrome(s)
	fmt.Println("a=", a)
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	check := func(l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l, r = l+1, r-1
		}
		return true
	}
	for l < r {
		if s[l] != s[r] {
			return check(l, r-1) || check(l+1, r)
		}
		l, r = l+1, r-1
	}
	return true
}
