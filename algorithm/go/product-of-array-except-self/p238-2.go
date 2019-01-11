package main

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	ret := make([]int, len(nums))

	// put productLeft in ret
	ret[0] = 1
	for i := 1; i < len(nums); i++ {
		ret[i] = ret[i-1] * nums[i-1]
	}

	productRight := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ret[i] = ret[i] * productRight
		productRight = productRight * nums[i]
	}
	return ret
}
