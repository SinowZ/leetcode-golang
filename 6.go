package main

import (
	"fmt"
)

func main() {
	ret := convert("LEETCODEISHIRING", 3)
	fmt.Println("ret=", ret)
}

func convert(s string, numRows int) string {
	if len(s) <= 2 || numRows == 1 {
		return s
	}
	var res string
	var tmp = make([]string, numRows)
	curPos := 0
	shouldTurn := -1
	for _, val := range s {
		tmp[curPos] += string(val)
		if curPos == 0 || curPos == numRows-1 {
			shouldTurn = -shouldTurn
		}
		curPos += shouldTurn
		fmt.Println("curPos=", curPos)
	}
	for _, val := range tmp {
		res += val
	}
	return res
}
