package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := findComplement(5)
	fmt.Println("a=", a)
}

func findComplement(num int) int {
	mc := strconv.FormatInt(int64(num), 2)

	bushu := ""
	for n := range mc {
		if mc[n] == 49 {
			bushu += "0"
		} else if mc[n] == 48 {
			bushu += "1"
		}
	}
	bushu = strings.TrimLeft(bushu, "0")

	ret := 0
	l := len(bushu)
	for i := l - 1; i >= 0; i-- {
		ret += (int(bushu[l-i-1]) & 0xf) << uint8(i)
	}

	return ret
}
