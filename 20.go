package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	str := isValid("{[]}")
	fmt.Println("------->", str)
}

func isValid(s string) bool {
	ret := s
	beforelength := math.MaxInt64
	afterlength := math.MaxInt64
	for beforelength != afterlength || (beforelength == math.MaxInt64) {
		beforelength = len(ret)
		ret = strings.ReplaceAll(ret, "{}", "")
		ret = strings.ReplaceAll(ret, "()", "")
		ret = strings.ReplaceAll(ret, "[]", "")
		afterlength = len(ret)
	}
	if afterlength == 0 {
		return true
	}
	return false
}
