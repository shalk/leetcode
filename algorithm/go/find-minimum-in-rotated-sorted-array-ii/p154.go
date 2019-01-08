package main

func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	mid := 0
	for l < r {
		if nums[l] < nums[r] {
			return nums[l]
		}
		mid = (l + r) >> 1
		if nums[mid] > nums[r] {
			l = mid + 1
		} else if nums[mid] < nums[r] {
			r = mid
		} else {
			r--
		}
	}
	return nums[l]
}
