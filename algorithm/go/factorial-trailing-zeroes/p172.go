package main

func trailingZeroes(n int) int {
	zero := 0
	divide := 5
	for n >= divide {
		zero = zero + (n / divide)
		divide = divide * 5
	}
	return zero
}
