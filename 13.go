package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ret := romanToInt("MCMXCIV")
	fmt.Println("ret= ", ret)
}

func romanToInt(s string) int {
	ret := 0
	s = strings.ReplaceAll(s, "IV", "4 ")
	s = strings.ReplaceAll(s, "IX", "9 ")
	s = strings.ReplaceAll(s, "XL", "40 ")
	s = strings.ReplaceAll(s, "XC", "90 ")
	s = strings.ReplaceAll(s, "CD", "400 ")
	s = strings.ReplaceAll(s, "CM", "900 ")
	s = strings.ReplaceAll(s, "I", "1 ")
	s = strings.ReplaceAll(s, "V", "5 ")
	s = strings.ReplaceAll(s, "X", "10 ")
	s = strings.ReplaceAll(s, "L", "50 ")
	s = strings.ReplaceAll(s, "C", "100 ")
	s = strings.ReplaceAll(s, "D", "500 ")
	s = strings.ReplaceAll(s, "M", "1000 ")
	s = strings.Trim(s, " ")
	array := strings.Split(s, " ")
	for v := range array {
		t, _ := strconv.Atoi(array[v])
		ret += t
	}
	return ret
}
