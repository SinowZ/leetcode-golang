package main

import (
	"fmt"
	"math"
)

func main() {
	//words := []string{"flower", "flow", "flight"}
	words := []string{"aa", "a"}
	//words := []string{"aca", "cba"}
	ret := longestCommonPrefix(words)
	fmt.Println("ret= ", ret)
}

func longestCommonPrefix(strs []string) string {
	ret := ""
	mc := map[int]string{}
	if len(strs) == 1 {
		ret = strs[0]
	}

	length := math.MaxInt64
	for _, str := range strs {
		if length > len(str) {
			length = len(str)
		}
	}
	if len(strs) > 1 {
		for i := 0; i <= length-1; i++ {
			for _, str := range strs {
				_, ok := mc[i]
				if !ok {
					mc[i] = string(str[i])
				} else {
					if mc[i] != string(str[i]) {
						delete(mc, i)
						break
					}
				}
			}

			_, ok := mc[i]
			if !ok {
				break
			}
			ret += mc[i]
		}
	}

	return ret
}
