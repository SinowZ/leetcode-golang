package main

import "fmt"

func main() {
	words := []string{"hello", "world", "leetcode"}
	chars := "welldonehoneyr"
	a := countCharacters(words, chars)
	fmt.Println("------>", a)
}

func countCharacters(words []string, chars string) int {
	TotalLength := 0

	for _, word := range words {
		TotalLength += len(word)

		mc := map[string]int{}
		for _, s := range chars {
			if _, ok := mc[string(s)]; !ok {
				mc[string(s)] = 0
			} else {
				mc[string(s)]++
			}
		}

		NotAwordLength := 0
		for _, w := range word {
			if _, ok := mc[string(w)]; !ok {
				NotAwordLength = len(word)
				break
			} else {
				if mc[string(w)] > 0 {
					mc[string(w)]--
				} else {
					delete(mc, string(w))
				}
			}
		}
		TotalLength -= NotAwordLength
	}
	return TotalLength
}
