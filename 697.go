package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 2, 3, 1, 4, 3, 3, 3}
	a := findShortestSubArray(nums)
	fmt.Println("a=", a)
}

func findShortestSubArray(nums []int) int {
	mc, mb := map[int]int{}, map[int]int{}
	d, out := 0, len(nums)
	for i, n := range nums {
		fmt.Println("i------->", i)
		/*
			mc[n]= 1 1
			mc[n]= 2 1
			mc[n]= 2 2
			mc[n]= 3 1
			mc[n]= 1 2
			mc[n]= 4 1
			mc[n]= 2 3
		*/
		mc[n]++
		if _, ok := mb[n]; !ok {
			mb[n] = i
			/*
				mb[n]= 1 0
				mb[n]= 2 1
				mb[n]= 3 3
				mb[n]= 4 5
			*/
		}
		if mc[n] == d {
			/*
				mc[n]= 2 1
				mc[n]= 1 2
			*/

			l := i - mb[n] + 1
			fmt.Println("i, n, mb[n], mc[n]----->", i, n, mb[n], mc[n], l)
			//第二轮 i=1, n=2, mb[n]=1, mc[n]=1， l=1
			//第五轮 i=4, n=1, mb[n]=0, mc[n]=2， l=5
			if l < out {
				out = l
			}
		} else if mc[n] > d {
			out, d = i-mb[n]+1, mc[n]
			fmt.Println("i-n-out-d--->", i, n, out, d)
			//第一轮 i=0, n=1, out=1, d=mc[n]=1
			//第三轮 i=2, n=2, out=2, d=mc[n]=2
			//第七轮 i=6, n=2, out=3, d=mc[n]=3
		}
	}
	return out
}
