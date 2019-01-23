package main

import "fmt"

func countPrimes(n int) int {
	count := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

func isPrime(i int) bool {
	if i < 2 {
		return false
	}
	if i == 2 {
		return true
	}
	if i%2 == 0 {
		return false
	}
	for k := 3; k <= int(math.Sqrt(float64(i))); k += 2 {
		if i%k == 0 {
			return false
		}
	}
	return true
}
