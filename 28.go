package main

import (
	"fmt"
)

func main() {
	ret := strStr("mississippi", "sipp")
	fmt.Println("ret= ", ret)
}

func strStr(haystack string, needle string) int {
	// ret := -1
	// if needle == "" || needle == haystack {
	// 	return 0
	// }
	// if len(needle) > len(haystack) {
	// 	return ret
	// }
	// isThis := false
	// for i := 0; i <= len(haystack)-len(needle); i++ {
	// 	if haystack[i] == needle[0] {
	// 		for j := 0; j < len(needle); j++ {
	// 			isThis = true
	// 			if needle[j] != haystack[i+j] {
	// 				isThis = false
	// 				break
	// 			}
	// 		}
	// 		if isThis {
	// 			return i
	// 		}
	// 	}
	// }
	// return ret

	if len(needle) == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
