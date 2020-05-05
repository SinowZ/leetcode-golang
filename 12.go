package main

import (
	"fmt"
)

func main() {
	ret := intToRoman(1994)
	fmt.Println("ret= ", ret)
}

func intToRoman(num int) string {
	charMap := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	intMap := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

	if num == 0 {
		return ""
	}
	index := len(charMap) - 1
	for ; index >= 0; index-- {
		if num-intMap[index] >= 0 {
			fmt.Println("index---->", index, num, intMap[index])
			break
		}
	}
	return charMap[index] + intToRoman(num-intMap[index])
}
