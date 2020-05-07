package main

import (
	"fmt"
)

func main() {
	a := generateTheString(10)
	fmt.Println("a=", a)
}

func generateTheString(n int) string {
	ret := ""
	if n%2 == 0 {
		ret = "a"
		for i := 1; i < n; i++ {
			ret += "b"
		}
	} else {
		for i := 0; i < n; i++ {
			ret += "b"
		}
	}
	return ret
}
