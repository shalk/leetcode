package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	if m == 0 {
		return 0
	}
	if m == 1 {
		return triangle[0][0]
	}
	// 0 .. m-1
	for i := m - 2; i >= 0; i-- {
		for k := 0; k <= i; k++ {
			if triangle[i+1][k] < triangle[i+1][k+1] {
				triangle[i][k] += triangle[i+1][k]
			} else {
				triangle[i][k] += triangle[i+1][k+1]
			}
		}
	}
	return triangle[0][0]
}
