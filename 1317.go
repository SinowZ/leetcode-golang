package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := getNoZeroIntegers(1010)
	fmt.Println("a=", a)
}

func getNoZeroIntegers(n int) []int {
	result := []int{}
	for i := 1; i < n; i++ {
		a := strconv.Itoa(i)
		b := strconv.Itoa(n - i)

		if strings.Index(a, "0") == -1 && strings.Index(b, "0") == -1 {
			result = append(result, i)
			result = append(result, n-i)
			break
		}
	}
	return result
}
