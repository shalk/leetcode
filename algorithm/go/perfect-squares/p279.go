package main

func numSquares(n int) int {

	if n == 1 {
		return 1
	}
	f := make([]int, n)
	f[0] = 1
	for i := 2; i <= n; i++ {
		f[i-1] = i
		for j := 1; j*j <= i; j++ {
			if j*j == i {
				f[i-1] = 1
			} else {
				f[i-1] = min(f[i-1], f[i-j*j-1]+1)
			}
		}
	}
	return f[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
