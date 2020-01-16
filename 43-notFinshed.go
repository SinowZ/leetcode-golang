package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := multiply("498828660196", "840477629533")
	fmt.Println("419254329864656431168468=", a)
}

func multiply(num1 string, num2 string) string {
	int1, _ := strconv.ParseFloat(num1, 64)
	int2, _ := strconv.ParseFloat(num2, 64)
	return strconv.FormatFloat(int1*int2, 'f', 0, 64)
}
