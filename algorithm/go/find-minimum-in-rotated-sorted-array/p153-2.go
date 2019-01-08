package main

import "fmt"

func findMin(nums []int) int {
	min := math.MaxInt32
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}
