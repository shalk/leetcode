package main

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	mid := 0
	for l < r {
		mid = (l + r) >> 1
		if nums[l] <= nums[r] {
			return nums[l]
		}
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}
