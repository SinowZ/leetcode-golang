package main

import (
	"fmt"
)

func main() {
	ret := longestPalindrome("abcba")
	fmt.Println("ret=", ret)
}

func longestPalindrome(s string) string {
	if len(s) == 1 || len(s) == 0 {
		return s
	} else if len(s) == 2 {
		if s[0:1] == s[1:2] {
			return s
		}
		return s[0:1]
	}
	ret := s[0:1]
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= i; j-- {
			if s[i] == s[j] && i != j {
				if len(s[i:j+1]) > len(ret) {
					isPalind := true
					for k := 0; k < len(s[i:j+1])-1; k++ {
						if s[i : j+1][k] != s[i : j+1][len(s[i:j+1])-1-k] {
							isPalind = false
							break
						}
					}
					if isPalind {
						ret = s[i : j+1]
					}
				}
			}
		}
	}
	return ret
}
