package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || matrix[0] == nil || len(matrix[0]) == 0 {
		return 0
	}
	// matrix = [][]byte{[]byte{'1'}}
	m := len(matrix)
	n := len(matrix[0])
	rightLen := make([][]int, m)
	downLen := make([][]int, m)
	for i := 0; i < m; i++ {
		rightLen[i] = make([]int, n)
		downLen[i] = make([]int, n)
	}
	// calc rightLen and downLen
	for i := 0; i < m; i++ {
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				if j == n-1 {
					rightLen[i][j] = 1
				} else {
					rightLen[i][j] = rightLen[i][j+1] + 1
				}
			}
		}
	}
	for j := 0; j < n; j++ {
		for i := m - 1; i >= 0; i-- {
			if matrix[i][j] == '1' {
				if i == m-1 {
					downLen[i][j] = 1
				} else {
					downLen[i][j] = downLen[i+1][j] + 1
				}
			}
		}
	}
	// fmt.Println(downLen)
	// fmt.Println(rightLen)
	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			minRight := rightLen[i][j]
			for k := i; k < i+downLen[i][j]; k++ {
				if rightLen[k][j] < minRight {
					minRight = rightLen[k][j]
				}
				if (k-i+1)*minRight > max {
					max = (k - i + 1) * minRight
				}
			}
		}
	}
	return max

}
