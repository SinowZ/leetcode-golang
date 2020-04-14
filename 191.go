package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "00000000000000000000000000001011"
	fmt.Println(s)
	u64, _ := strconv.ParseUint(s, 2, 32)
	u32 := uint32(u64)
	a := hammingWeight(u32)
	fmt.Println("a=", a)
}

func hammingWeight(num uint32) int {
	var out uint32
	for num != 0 {
		out += num & 1
		num >>= 1
	}
	return int(out)
}
