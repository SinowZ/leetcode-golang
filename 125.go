package main

import (
	"fmt"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	a := isPalindrome(s)
	fmt.Println("a=", a)
}

func isPalindrome(s string) bool {
	result := []rune(s)
	newStr := ""
	for _, v := range result {
		if v >= 65 && v <= 90 {
			v += 32
		}
		if (v >= 48 && v <= 57) || (v >= 97 && v <= 122) {
			newStr += string(v)
		}
	}
	fmt.Println("newStr=", newStr)
	for i, _ := range newStr {
		if newStr[i] != newStr[len(newStr)-i-1] {
			return false
		}
	}
	return true
}
