package main

import (
	"fmt"
)

func main() {
	a := missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
	b := missingNumber([]int{0})
	c := missingNumber([]int{1})
	d := missingNumber([]int{0, 1})
	e := missingNumber([]int{1, 2})
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)
	fmt.Println("e=", e)
}

func missingNumber(nums []int) int {
	//好方法
	// m := make([]byte, len(nums)+1)
	// for _, v := range nums {
	//     m[v] = 1
	// }
	// for k, v := range m {
	//     if v != 1 {
	//         return k
	//     }
	// }
	// return -1

	//笨方法
	m := map[int]bool{}
	min := nums[0]
	max := nums[0]
	for k := range nums {
		if nums[k] < min {
			min = nums[k]
		}
		if nums[k] > max {
			max = nums[k]
		}
		m[nums[k]] = true
	}

	for i := min; i < max; i++ {
		_, ok := m[i]
		if ok == false {
			return i
		}
	}

	if min != 0 {
		return min - 1
	}
	return max + 1
}
