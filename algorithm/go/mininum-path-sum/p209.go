package main

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ret := len(nums) + 1
	start, end := 0, 0
	sum := nums[0]
	for end < len(nums) {
		if sum >= s {
			if end-start+1 < ret {
				ret = end - start + 1
			}
			if start < end {
				sum -= nums[start]
				start++
			} else {
				return 1
			}
		} else { // sum < s
			end++
			if end < len(nums) {
				sum += nums[end]
			}
		}
	}
	if ret == len(nums)+1 {
		return 0
	} else {
		return ret
	}
}
