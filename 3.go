package main

import "fmt"

func main() {
	orignal := "abcdabcdabcbb"
	lastKeyIndex := make(map[rune]int)
	start := 0
	maxLenth := 0

	for index, c := range orignal {

		if lastI, ok := lastKeyIndex[c]; ok && lastI >= start {
			start = lastI + 1
		}
		if newLen := index - start + 1; newLen > maxLenth {
			maxLenth = newLen
		}
		lastKeyIndex[c] = index
	}

	fmt.Println("------->", maxLenth)
}
