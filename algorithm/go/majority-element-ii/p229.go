package main

import "fmt"

func majorityElement(nums []int) []int {
	ret := make([]int, 0)
	if len(nums) == 0 {
		return ret
	}
	num1, num2 := nums[0], nums[0]
	count1, count2 := 0, 0
	for _, num := range nums {
		if num == num1 {
			count1++
		} else if num == num2 {
			count2++
		} else if count1 == 0 {
			num1 = num
			count1 = 1
		} else if count2 == 0 {
			num2 = num
			count2 = 1
		} else {
			count1--
			count2--
		}
	}
	count1, count2 = 0, 0
	for _, num := range nums {
		if num == num1 {
			count1++
		} else if num == num2 {
			count2++
		}
	}

	if count1 > len(nums)/3 {
		ret = append(ret, num1)
	}
	if num2 != num1 && count2 > len(nums)/3 {
		ret = append(ret, num2)
	}
	return ret
}
