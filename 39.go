package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 3, 5, 6}
	target := 8
	a := combinationSum(nums, target)
	fmt.Println("a=", a)
}

func combinationSum(candidates []int, target int) [][]int {
	var (
		ans    [][]int
		res    []int
		length = len(candidates)
		dfs    func(int)
		last   int
	)

	dfs = func(target int) {
		if target < 0 {
			return
		}
		if target == 0 {
			tmp := make([]int, len(res))
			copy(tmp, res)
			ans = append(ans, tmp)
		}
		for i := last; i < length; i++ {
			if target-candidates[i] >= 0 {
				last = i
				res = append(res, candidates[i])
				dfs(target - candidates[i])
				res = res[:len(res)-1]
				last = i
			}
		}
	}

	sort.Ints(candidates)
	dfs(target)
	return ans
}
