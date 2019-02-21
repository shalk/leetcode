package main

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	left := 1
	right := x
	// first last k  k*k < x
	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if mid*mid > x {
			right = mid - 1
		} else if mid*mid < x {
			left = mid + 1
		} else {
			return mid
		}
	}
	return right
}
