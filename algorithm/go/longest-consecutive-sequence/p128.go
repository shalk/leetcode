package main

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	h := make(map[int]int)

	for _, num := range nums {
		if _, ok := h[num]; ok {
			continue
		}
		h[num] = 1
		if num == math.MinInt32 {
			len1, ok1 := h[num+1]
			if ok1 {
				r := num + len1
				h[r] = len1 + 1
				h[num] = len1 + 1
			}
		} else if num == math.MaxInt32 {
			len2, ok2 := h[num-1]
			if ok2 {
				l := num + len2
				h[l] = len2 + 1
				h[num] = len2 + 1
			}
		} else {
			len1, ok1 := h[num+1]
			len2, ok2 := h[num-1]
			if ok1 && ok2 {
				r := num + len1
				l := num - len2
				h[r] = len1 + len2 + 1
				h[l] = len1 + len2 + 1
			} else if ok1 {
				r := num + len1
				h[r] = len1 + 1
				h[num] = len1 + 1
			} else if ok2 {
				l := num - len2
				h[l] = len2 + 1
				h[num] = len2 + 1
			} else {

			}
		}

	}

	ret := 1
	for _, len1 := range h {
		if len1 > ret {
			ret = len1
		}
	}
	return ret
}
