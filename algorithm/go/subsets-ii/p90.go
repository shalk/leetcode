package main

import "fmt"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0)
	sub := make([]int, 0)
	helper(&ret, sub, 0, nums)
	return ret
}

func helper(ret *[][]int, sub []int, pos int, nums []int) {
	if pos == len(nums) {
		*ret = append(*ret, append([]int{}, sub...))
		return
	}
	count := 1
	for i := pos + 1; i < len(nums); i++ {
		if nums[i] == nums[pos] {
			count++
		}
	}
	nextPos := pos + count

	for i := 0; i <= count; i++ {
		for k := 0; k < i; k++ {
			sub = append(sub, nums[pos])
		}
		helper(ret, sub, nextPos, nums)
		sub = sub[:len(sub)-i]
	}

}
