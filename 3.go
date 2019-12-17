package main

import "fmt"

func main() {
	orignal := "abcdabcdabcbb"
	/*
		var maxStr []string
		var tempStr string
		orignals := []string{orignal}
		for i := 0; i < len(orignals); i++ {
			str := orignals[i]

			if strings.Contains(tempStr, str) {

				if len(tempStr) > len(maxStr) {
					maxStr = []string{tempStr}
				}

				reIndex := strings.Index(tempStr, str)
				tempStr = tempStr[reIndex+1 : len(tempStr)-1]

			}
			tempStr = tempStr + str
		}

		if len(tempStr) > len(maxStr) {
			maxStr = []string{tempStr}
		}

		fmt.Println("------------->", len(maxStr))
	*/

	lastKeyIndex := make(map[rune]int)
	start := 0
	maxLenth := 0

	for index, c := range orignal { // 下标， 值

		if lastI, ok := lastKeyIndex[c]; ok && lastI >= start {
			start = lastI + 1
			fmt.Println("start------->", index, start, maxLenth)
		}
		if newLen := index - start + 1; newLen > maxLenth {
			fmt.Println("newLen------->", index, start, newLen)
			maxLenth = newLen
		}
		lastKeyIndex[c] = index
	}

	//fmt.Println("------->", maxLenth)
}
