package main

import (
	"fmt"
)

func main() {
	fmt.Println("-1A1B------>", getHint("1123", "0111"))
	fmt.Println("-1A3B------>", getHint("1807", "7810"))
	fmt.Println("-1A0B------>", getHint("11", "01"))
	fmt.Println("-1A0B------>", getHint("11", "10"))
	fmt.Println("-0A4B------>", getHint("1122", "2211"))
}

func getHint(secret string, guess string) string {
	m := [256]int{}
	nb, nc := 0, 0
	for i, s := range []byte(secret) {
		if s == guess[i] {
			nb++
		} else {
			if m[s] < 0 {
				nc++
			}
			if m[guess[i]] > 0 {
				nc++
			}
			m[s]++
			m[guess[i]]--
		}
	}
	return fmt.Sprintf("%dA%dB", nb, nc)
}
