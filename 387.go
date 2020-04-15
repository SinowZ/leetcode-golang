package main

import (
	"fmt"
)

func main() {
	s := "loveleetcode"
	a := firstUniqChar(s)
	fmt.Println("a=", a)
}

func firstUniqChar(s string) int {
	mc := map[string]int{}

	for _, v := range []rune(s) {
		if _, ok := mc[string(v)]; !ok {
			mc[string(v)] = 0
		} else {
			mc[string(v)]++
		}
	}

	for i, v := range []rune(s) {
		if _, ok := mc[string(v)]; ok && mc[string(v)] == 0 {
			return i
		}
	}

	return -1
}
