package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	productLeft := make([]int, len(nums))
	productRight := make([]int, len(nums))
	productLeft[0] = 1
	for i := 1; i < len(nums); i++ {
		productLeft[i] = productLeft[i-1] * nums[i-1]
	}
	productRight[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		productRight[i] = productRight[i+1] * nums[i+1]
	}
	ret := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ret[i] = productRight[i] * productLeft[i]
	}
	return ret
}
