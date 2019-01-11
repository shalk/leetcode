package main

import "fmt"

func findDuplicate(nums []int) int {
	l := 1
	r := len(nums)
	for l < r {
		mid := (l + r) / 2
		// count nums in range [1.. mid]
		count := 0
		for _, num := range nums {
			if num >= l && num <= mid {
				count++
			}
		}
		if count > (mid-l)+1 {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
