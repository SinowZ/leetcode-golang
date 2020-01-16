package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := addBinary("11", "1")
	fmt.Println("a=", a)
}

func addBinary(a string, b string) string {
	numsa := []int{}
	numsb := []int{}
	for _, v := range []rune(a) {
		n, _ := strconv.Atoi(string(v))
		numsa = append(numsa, n)
	}

	for _, v := range []rune(b) {
		n, _ := strconv.Atoi(string(v))
		numsb = append(numsb, n)
	}

	fmt.Println("numsa=", numsa)
	fmt.Println("numsb=", numsb)

	return a
}
