package main

import "fmt"

func maxProfit(prices []int) int {
	// greedy
	ret := 0
	buy := -1
	sell := -1
	for i := 0; i < len(prices); i++ {
		// find buy time
		// tomorrow price is cheaper
		for i+1 < len(prices) && prices[i] >= prices[i+1] {
			i++
		}
		buy = prices[i]
		if i+1 == len(prices) {
			break
		}
		// tomorrow price is more expensive
		for i+1 < len(prices) && prices[i] <= prices[i+1] {
			i++
		}
		sell = prices[i]
		ret = ret + sell - buy
	}
	return ret
}
