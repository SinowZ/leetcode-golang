package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	a := relativeSortArray(arr1, arr2)
	fmt.Println("a=", a)
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	ret := []int{}
	ret2 := []int{}
	var notBelong bool

	for _, i := range arr2 {
		for _, j := range arr1 {
			if i == j {
				ret = append(ret, i)
			}
		}
	}

	for _, i := range arr1 {
		notBelong = false
		for _, j := range arr2 {
			if i == j {
				notBelong = true
			}
		}
		if notBelong == false {
			ret2 = append(ret2, i)
		}
	}

	sort.Ints(ret2)

	return append(ret, ret2...)
}
