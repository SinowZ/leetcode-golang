package main

import (
	"fmt"
)

func main() {
	a := climbStairs(7)
	fmt.Println("a=", a)
}

func climbStairs(n int) int {

	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	t := make([]int, n)
	t[0], t[1] = 1, 2
	for i := 2; i < n; i++ {
		t[i] = t[i-1] + t[i-2]
	}
	return t[n-1]

	// str := ""
	// mc := map[string]int{}

	// for i := 0; i < n; i++ {
	// 	str += "1"
	// }
	// mc[str] = 0

	// for i := 0; i < len(str); i++ {
	// 	temp := str[i:]
	// 	prefix := str[:i]
	// 	if strings.Index(temp, "11") != -1 {
	// 		mc[prefix+strings.Replace(temp, "11", "2", 1)] = 0
	// 	}
	// }

	// for k, _ := range mc {
	// 	for j := 0; j < len(k); j++ {
	// 		temp1 := k[j:]
	// 		prefix1 := k[:j]
	// 		if strings.Index(temp1, "11") != -1 {
	// 			mc[prefix1+strings.Replace(temp1, "11", "2", 1)] = 0
	// 			fmt.Println("3=", prefix1+strings.Replace(temp1, "11", "2", 1))
	// 		}
	// 	}
	// }

	// fmt.Println("mc=", mc)

	// return len(mc)
}
