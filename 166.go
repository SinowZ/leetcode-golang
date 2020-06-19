package main

import (
	"fmt"
)

func main() {
	a := fractionToDecimal(1, 2)
	fmt.Println("a=", a)
}

func fractionToDecimal(numerator int, denominator int) string {
	n := float64(numerator) / float64(denominator)
	fmt.Println("--------->", n)
	//fmt.Sprintf("%.2f", value)

	return ""
}
