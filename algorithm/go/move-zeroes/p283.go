package main

import "fmt"

func moveZeroes(nums []int) {
	p := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if p != i {
				nums[p], nums[i] = nums[i], nums[p]
			}
			p++
		}
	}
}
