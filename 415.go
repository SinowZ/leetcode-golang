package main

import (
	"fmt"
)

func main() {

	a := addStrings("1", "9")
	fmt.Println("---------", a)
}

func addStrings(num1 string, num2 string) string {

	out := ""
	i, j := len(num1)-1, len(num2)-1
	var c byte = 0
	for i >= 0 || j >= 0 {
		var a, b byte = 0, 0
		if i >= 0 {
			a = num1[i] - '0'
			i--
		}
		if j >= 0 {
			b = num2[j] - '0'
			j--
		}
		s := a + b + c
		out = string('0'+s%10) + out
		c = s / 10
	}
	if c != 0 {
		return "1" + out
	}
	return out

	/* 笨方法
	ret := ""
	tenFlag := false
	prefix := ""
	lennum1 := 0

	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}
	prefix = num2[:len(num2)-len(num1)]
	lennum1 = len(num1)

	for len(num1) > 0 {

		lstr1 := num1[len(num1)-1:]
		lint1, _ := strconv.Atoi(lstr1)

		lstr2 := num2[len(num2)-1:]
		lint2, _ := strconv.Atoi(lstr2)
		sum := lint1 + lint2
		if tenFlag {
			sum += 1
		}

		if len(num1) != 1 {
			if sum <= 9 {
				ret = strconv.Itoa(sum) + ret
				tenFlag = false
			} else {
				ret = strconv.Itoa(sum%10) + ret
				tenFlag = true
			}
			num1 = num1[0 : len(num1)-1]
			num2 = num2[0 : len(num2)-1]
		} else {
			ret = strconv.Itoa(sum) + ret
			num1 = ""
			num2 = ""
		}
	}

	if lennum1 == len(ret) {
		ret = prefix + ret
	} else {
		firstNum, _ := strconv.Atoi(ret[:1])
		ret = ret[1:]

		tenFlag = false
		for len(prefix) > 0 {
			lstrp := prefix[len(prefix)-1:]
			lintp, _ := strconv.Atoi(lstrp)

			sum := lintp + firstNum
			if tenFlag {
				sum += 1
			}

			if sum <= 9 {
				ret = strconv.Itoa(sum) + ret
				tenFlag = false
			} else {
				ret = strconv.Itoa(sum%10) + ret
				tenFlag = true
			}
			prefix = prefix[0 : len(prefix)-1]
			firstNum = 0
		}
	}
	return ret */
}
