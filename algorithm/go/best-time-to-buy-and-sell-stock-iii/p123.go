package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	ret := 0
	//buy at i, sell at j ; then find maxProfit after j
	// before[i] is maxProfit at [0...i] for at most one transaction
	before := make([]int, len(prices))
	// after[i] is maxProfit at [i+1...len(prices)-1] for at most one transaction
	after := make([]int, len(prices))

	min1 := prices[0]
	before[0] = 0
	for i := 1; i < len(prices); i++ {
		before[i] = max(before[i-1], prices[i]-min1)
		min1 = min(min1, prices[i])
	}

	after[len(prices)-1] = 0
	after[len(prices)-2] = 0
	max1 := prices[len(prices)-1]
	for i := len(prices) - 3; i >= 0; i-- {
		after[i] = max(after[i+1], max1-prices[i+1])
		max1 = max(max1, prices[i+1])
	}

	for i := 0; i < len(prices); i++ {
		ret = max(ret, after[i]+before[i])
	}
	return ret
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
