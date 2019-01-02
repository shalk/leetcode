package main

import "fmt"

func maxProfit(prices []int) int {
	ret := 0
	if len(prices) <= 1 {
		return 0
	}
	max := prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		if max-prices[i] > ret {
			ret = max - prices[i]
		}
		if prices[i] > max {
			max = prices[i]
		}
	}
	return ret
}
