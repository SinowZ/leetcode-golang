package main

import (
	"fmt"
	"math"
)

func main() {
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"KFC", "Shogun", "Burger King"}
	num := findRestaurant(list1, list2)
	fmt.Println("num=", num)
}

func findRestaurant(list1 []string, list2 []string) []string {
	count := math.MaxInt64
	ret := []string{}

	mc := map[string]int{}
	md := map[string]int{}
	for i1, s1 := range list1 {
		mc[s1] = i1
	}
	for i2, v := range list2 {
		_, ok := mc[v]
		if ok {
			md[v] = mc[v] + i2
			if mc[v]+i2 < count {
				count = mc[v] + i2
			}
		}
	}

	for k, v := range md {
		if v == count {
			ret = append(ret, k)
		}
	}

	return ret
}
