package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "foobarfoobar"
	words := []string{"foo", "bar"}
	num := findSubstring(s, words)
	fmt.Println("-------->", num)
}

func findSubstring(s string, words []string) []int {
	results := []int{}
	dumpSingle := ""
	if len(words) == 0 {
		return results
	} else {
		for i := 0; i < len(words[0]); i++ {
			dumpSingle += "@"
		}

		if len(words) == 1 {
			for strings.Contains(s, words[0]) {
				index := strings.Index(s, words[0])
				results = append(results, index)
				s = strings.Replace(s, words[0], dumpSingle, 1)
			}
		} else {
			words2 := []string{}

			// for n, _ := range words {
			// 	words2 = append(words2, string(n+1))
			// }
			words2 = append(words2, "1")
			words2 = append(words2, "2")

			result := outOrder(words2)

			fmt.Println("-result------->", result)

			temp := ""
			for j, v := range result {
				temp = ""
				for i := 0; i < len(v); i++ {
					temp += words[int(v[i])-1]
					fmt.Println("-temp------->", temp)
				}
				index := strings.Index(s, temp)
				if index != -1 {
					containIndex := false
					for _, v := range results {
						if v == index {
							containIndex = true
						}
					}
					if !containIndex {
						results = append(results, index)
					}
					containIndex = false
					fmt.Println("-index------->", index)
				}
				j++
			}
			return results
		}
	}
	return results
}

func outOrder(trainsNums []string) []string {
	COUNT := len(trainsNums)
	//如果只有一个数，则直接返回
	if COUNT == 1 {
		return []string{trainsNums[0]}
	}
	//否则，将最后一个数插入到前面的排列数中的所有位置（递归）
	return insert(outOrder(trainsNums[:COUNT-1]), trainsNums[COUNT-1])
}

func insert(res []string, insertNum string) []string {
	//保存结果的slice
	result := make([]string, len(res)*(len(res[0])+1))
	index := 0
	for _, v := range res {
		for i := 0; i < len(v); i++ {
			//在v的每一个元素前面插入
			result[index] = v[:i] + insertNum + v[i:]
			index++
		}
		//在v最后面插入
		result[index] = v + insertNum
		index++
	}
	return result
}
