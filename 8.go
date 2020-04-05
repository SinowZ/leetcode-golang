package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ret := myAtoi("+-1")
	fmt.Println("ret-----> ", ret)
}

func myAtoi(str string) int {
	str = strings.TrimSpace(str)

	if str == "" {
		return 0
	}

	ret := ""
	neg := false

	if str[0] == 43 {
		str = str[1:]
	} else if str[0] == 45 {
		str = str[1:]
		neg = true
	}

	for i := range str {
		if str[i] >= 48 && str[i] <= 57 {
			ret += string(str[i])
		} else {
			break
		}
	}

	if neg {
		ret = "-" + ret
	}

	retint32, _ := strconv.ParseInt(ret, 10, 32)
	return int(retint32)
}
