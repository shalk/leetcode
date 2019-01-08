package main

import "fmt"

func majorityElement(nums []int) int {
	// n2 := len(nums)
	ret := nums[0]
	count := 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			count = 1
			ret = nums[i]
			continue
		}
		if nums[i] == ret {
			count++
		} else {
			count--
		}
	}
	return ret
}
