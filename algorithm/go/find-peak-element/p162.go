package main

import "fmt"

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}
	l, r := 1, len(nums)-2
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		} else if nums[mid] <= nums[mid-1] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}
