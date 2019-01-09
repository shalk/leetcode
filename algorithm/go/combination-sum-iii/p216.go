package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	ret := make([][]int, 0)
	helper(&ret, make([]int, 0), k, n)
	return ret
}

func helper(ret *[][]int, sub []int, k int, n int) {
	if len(sub) > k {
		return
	}
	if len(sub) == k {
		sum := 0
		for _, v := range sub {
			sum += v
		}
		if sum == n {
			*ret = append(*ret, append([]int{}, sub...))
		}
		return
	}
	if len(sub) == 0 {
		for i := 1; i <= 9; i++ {
			sub = append(sub, i)
			helper(ret, sub, k, n)
			sub = sub[:len(sub)-1]
		}
	} else {
		for i := sub[len(sub)-1] + 1; i <= 9; i++ {
			sub = append(sub, i)
			helper(ret, sub, k, n)
			sub = sub[:len(sub)-1]
		}
	}
}
