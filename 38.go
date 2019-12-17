package main

import (
	"fmt"
)

func main() {
	str := countAndSay(12)
	fmt.Println("------->", str)
}

func countAndSay(n int) string {
	newStr := ""
	if n == 1 {
		return "1"
	} else {
		lastStr := countAndSay(n - 1)
		cur := lastStr[0:1]
		curT := 0

		for i := 1; i < len(laststr); i++ {
			if laststr[i:i+1] == cur {
				curT++
			} else {
				new_str += str(cur_t) + cur
				cur = i
				cur_t = 1
			}
		}
		new_str += str(cur_t) + cur
	}

	return new_str

	/*
		str1 := "1"
		str2 := ""
		var left int
		var count int

		for i := 1; i < n; i++ {
			left = 0
			for left < len(str1) { // 0 < 1
				count = 0
				for left+count < len(str1) && str1[left:left+1] == str1[left+count:left+count+1] { //1==1
					count++
				}
				str2 += string(count) + string(str1[left:left+1]) //11
				left += count
			}
			str1 += str2
			str2 = ""
		}
		return str1*/
}
