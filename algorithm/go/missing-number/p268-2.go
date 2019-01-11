package main

func missingNumber(nums []int) int {
	for i := 0; i < len(nums); {
		if i != nums[i] && nums[i] < len(nums) {
			j := nums[i]
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}
	return len(nums)
}
