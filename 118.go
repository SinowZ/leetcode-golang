package main

import (
	"fmt"
)

func main() {
	a := generate(5)
	fmt.Println("a=", a)
}

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	if numRows == 1 {
		return [][]int{{1}}
	}

	var ret = make([][]int, numRows)
	ret[0] = []int{1}
	for i := 1; i < numRows; i++ {
		var tmp = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				tmp[j] = 1
				continue
			}
			tmp[j] = ret[i-1][j] + ret[i-1][j-1]
		}
		ret[i] = tmp
	}
	return ret
}
