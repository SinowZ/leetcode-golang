package main

import (
	"fmt"
)

func main() {
	a := countPrimes(1500000)
	fmt.Println("a=", a)
}

func countPrimes(n int) int {
	isPrime := make([]bool, n)

	for i := 2; i < n; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if !isPrime[i] {
			continue
		}
		for j := i * i; j < n; j += i {
			isPrime[j] = false
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}

	return count
}
