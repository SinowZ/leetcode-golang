package main

import (
	"fmt"
)

func main() {
	val := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	ret := maxArea(val)
	fmt.Println("ret= ", ret)
}

func maxArea(height []int) int {
	// ret := 0
	// for i := 0; i < len(height); i++ {
	// 	for j := len(height) - 1; j > i; j-- {
	// 		temp := height[i]
	// 		if height[j] < temp {
	// 			temp = height[j]
	// 		}
	// 		temp = temp * (j - i)
	// 		if temp > ret {
	// 			ret = temp
	// 		}
	// 	}
	// }

	ret := 0
	i := 0
	j := len(height) - 1
	for i < j {
		temp := height[i]
		if height[j] < temp {
			temp = height[j]
		}
		temp = temp * (j - i)
		if temp > ret {
			ret = temp
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return ret
}
