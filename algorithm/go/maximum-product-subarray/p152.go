package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func maxProduct(nums []int) int {
	ret := nums[0]
	start := -1
	end := -1
	for i := 0; i < len(nums); i++ {
		if ret < nums[i] {
			ret = nums[i]
		}
		if nums[i] == 0 {
			continue
		}
		if start == -1 {
			start = i
		}
		if i == len(nums)-1 || nums[i+1] == 0 {
			end = i
			cur := helper(nums, start, end)
			if cur > ret {
				ret = cur
			}
			start = -1
			end = -1
		}
	}
	return ret
}

func helper(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}
	ret := 1
	for i := start; i <= end; i++ {
		ret = ret * nums[i]
	}
	if ret > 0 {
		return ret
	}
	// remove left part
	ret1 := ret
	for i := start; i <= end; i++ {
		if nums[i] > 0 {
			ret1 = ret1 / nums[i]
		} else {
			ret1 = ret1 / nums[i]
			break
		}
	}

	ret2 := ret
	for i := end; i >= start; i-- {
		if nums[i] > 0 {
			ret2 = ret2 / nums[i]
		} else {
			ret2 = ret2 / nums[i]
			break
		}
	}
	if ret1 > ret {
		ret = ret1
	}
	if ret2 > ret {
		ret = ret2
	}
	return ret
}
