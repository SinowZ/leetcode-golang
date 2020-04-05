package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := fizzBuzz(15)
	fmt.Println("a=", a)
}

func fizzBuzz(n int) []string {

	result := []string{}

	for i := 1; i < n+1; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
			fmt.Println("---------", strconv.Itoa(i))
		}
	}

	return result
}
