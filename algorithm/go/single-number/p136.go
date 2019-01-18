package main

import "fmt"

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	ret := 0
	for _, num := range nums {
		ret = ret ^ num
	}
	return ret
}
