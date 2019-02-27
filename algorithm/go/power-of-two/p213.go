package main

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	for n > 1 {
		if n&1 == 0 {
			n = n >> 1
		} else {
			return false
		}
	}
	return true
}
