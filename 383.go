package main

import (
	"fmt"
)

func main() {
	a := canConstruct("aa", "aab")
	fmt.Println("a=", a)
}

func canConstruct(ransomNote string, magazine string) bool {
	mc := map[string]int{}

	for _, v := range []rune(magazine) {
		if _, ok := mc[string(v)]; !ok {
			mc[string(v)] = 0
		} else {
			mc[string(v)]++
		}
	}

	for _, v := range []rune(ransomNote) {
		if _, ok := mc[string(v)]; !ok {
			return false
		} else if mc[string(v)] > 0 {
			mc[string(v)]--
		} else if mc[string(v)] == 0 {
			delete(mc, string(v))
		}
	}

	return true
}
