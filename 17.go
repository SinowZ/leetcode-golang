package main

import (
	"fmt"
)

func main() {
	str := letterCombinations("23569")
	fmt.Println("------->", str)
}

func letterCombinations(digits string) []string {
	var array = [9][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}, {"j", "k", "l"},
		{"m", "n", "o"}, {"p", "q", "r", "s"}, {"t", "u", "v"}, {"w", "x", "y", "z"}}

	result := []string{}

	if len(digits) == 0 {
		return result
	}

	bs := array[digits[0]-'2']

	if 1 == len(digits) {
		for _, b := range bs {
			result = append(result, b)
		}
		return result
	}

	subs := letterCombinations(digits[1:])
	for _, b := range bs {
		for _, s := range subs {
			result = append(result, b+s)
		}
	}
	return result
}
