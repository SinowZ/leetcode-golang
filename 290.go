package main

import (
	"fmt"
	"strings"
)

func main() {
	a := wordPattern("abba", "dog dog cat dog")
	fmt.Println("a=", a)
}

func wordPattern(pattern string, str string) bool {

	strs := strings.Split(str, " ")
	if len(pattern) != len(strs) {
		return false
	}

	m := map[byte]string{}
	p := map[byte]int{}
	s := map[string]int{}

	for k := range pattern {
		p[pattern[k]]++
		_, ok := m[pattern[k]]
		if ok {
			if m[pattern[k]] != strs[k] {
				return false
			}
		} else {
			m[pattern[k]] = strs[k]
		}
	}

	for k := range strs {
		s[strs[k]] = k
	}

	return len(p) == len(s)
}
