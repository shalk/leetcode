package main

import "fmt"

func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}

	k = k % len(nums)
	reverse(nums, 0, len(nums)-k-1)         // len is len(nums) - k
	reverse(nums, len(nums)-k, len(nums)-1) // len is k
	reverse(nums, 0, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for i, j := start, end; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
