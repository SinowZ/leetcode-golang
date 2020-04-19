package main

import (
	"fmt"
	"strings"
)

func main() {
	a := lengthOfLastWord("b   a    ")
	fmt.Println("a=", a)
}

func lengthOfLastWord(s string) int {
	ret := 0

	s = strings.Trim(s, " ")
	if strings.Index(s, " ") == -1 {
		return len(s)
	} else {

		mc := strings.Split(s, " ")

		if len(mc) >= 2 {
			last := mc[len(mc)-1]
			return len(last)
		}
	}

	return ret
}
