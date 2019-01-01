package main

import "fmt"

func generate(numRows int) [][]int {
	ret := make([][]int, numRows)
	if numRows == 0 {
		return ret
	}
	ret[0] = []int{1}
	if numRows == 1 {
		return ret
	}
	ret[1] = []int{1, 1}
	if numRows == 2 {
		return ret
	}
	for i := 2; i < numRows; i++ {
		// nums[i][0] =
		ret[i] = make([]int, i+1)
		ret[i][0] = 1
		ret[i][i] = 1
		for k := 1; k < i; k++ {
			ret[i][k] = ret[i-1][k-1] + ret[i-1][k]
		}
	}
	return ret

}
