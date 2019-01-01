package main

import "fmt"

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}

	ret := make([]int, rowIndex+1)
	ret[0] = 1
	ret[1] = 1
	for i := 2; i < rowIndex+1; i++ {
		// fmt.Println(ret)
		ret[0] = 1
		ret[i] = 1
		for k := i - 1; k >= 1; k-- {
			ret[k] = ret[k-1] + ret[k]
		}
	}
	return ret
}
