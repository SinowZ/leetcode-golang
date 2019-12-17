package main

import (
	"fmt"
)

func main() {
	prices := []int{7, 6, 4, 3, 1}
	a := maxProfit(prices)
	fmt.Println("a=", a)
}

func maxProfit(prices []int) int {
	max := 0
	for i, v := range prices {
		for j, w := range prices {
			if j > i && w > v && w-v > max {
				max = w - v
			}
			j++
		}
		i++
	}
	return max
}
