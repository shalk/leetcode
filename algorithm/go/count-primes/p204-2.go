package main

//https://zh.wikipedia.org/wiki/%E5%9F%83%E6%8B%89%E6%89%98%E6%96%AF%E7%89%B9%E5%B0%BC%E7%AD%9B%E6%B3%95
func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	notPrime := make([]bool, n)
	// notPrime[2] = true
	count := 0
	for i := 2; i < n; i++ {
		if !notPrime[i] {
			count++
			for j := i * i; j < n; j += i {
				notPrime[j] = true
			}
		}
	}
	return count
}
