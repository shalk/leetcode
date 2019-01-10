package main

import "fmt"

func summaryRanges(nums []int) []string {
	ret := make([]string, 0)

	size := 0
	for i := 0; i < len(nums); i++ {
		// check if nums[i] in range before
		if size == 0 || nums[i] == nums[i-1]+1 {
			size++
		} else {
			if size == 1 {
				ret = append(ret, fmt.Sprintf("%d", nums[i-1]))
			} else {
				ret = append(ret, fmt.Sprintf("%d->%d", nums[i-size], nums[i-1]))
			}
			size = 1
		}
	}
	if size == 0 {

	} else if size == 1 {
		ret = append(ret, fmt.Sprintf("%d", nums[len(nums)-1]))
	} else {
		ret = append(ret, fmt.Sprintf("%d->%d", nums[len(nums)-size], nums[len(nums)-1]))
	}
	return ret
}
