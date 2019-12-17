package main

import "fmt"

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	//s := []byte{}

	reverseString(s)
	fmt.Println("------>", s)
}

func reverseString(s []byte) {
	fmt.Println("------>", s)
	if len(s) != 0 && len(s) != 1 {
		for i := 0; i >= len(s)-1; i++ {
			s[i] = s[len(s)-1-i]
		}
	}
}
