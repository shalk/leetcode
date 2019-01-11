package main

func missingNumber(nums []int) int {
	expect := 0
	for i := 0; i < len(nums)+1; i++ {
		expect += i
	}
	actual := 0
	for _, num := range nums {
		actual += num
	}
	return expect - actual
}
