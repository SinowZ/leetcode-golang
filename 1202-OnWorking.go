package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "dcab"
	pairs := [][]int{{5, 6}, {0, 3}, {1, 2}, {0, 2}, {4, 6}}
	a := smallestStringWithSwaps(s, pairs)
	fmt.Println("a=", a)
}

//"dcab"
func smallestStringWithSwaps(s string, pairs [][]int) string {
	result := []rune(s)
	orders := map[int][]int{}
	for i, v := range pairs {
		if i == 0 {
			orders[0] = v
		} else {
			for k, _ := range orders {
				if v[0] >= orders[k][0] && v[1] <= orders[k][len(orders[k])-1] {
					orders[k] = append(orders[k], v[0])
					orders[k] = append(orders[k], v[1])
				} else {
					orders[i] = v
				}
			}
		}
		sort.Ints(orders[0])

	}
	for key, value := range orders {
		orders[key] = removeDup(value)
	}

	fmt.Println("---->", orders)

	return string(result)
}

func removeDup(array []int) []int {
	out := []int{}
	mp1 := map[int]int{}
	for _, v := range array {
		mp1[v]++
	}
	for k, _ := range mp1 {
		out = append(out, k)
	}
	return out
}

//if result[v[1]] < result[v[0]] {
//	result[v[0]], result[v[1]] = result[v[1]], result[v[0]]
//}
//fmt.Println("after---->", string(result), v[0], v[1])*/
