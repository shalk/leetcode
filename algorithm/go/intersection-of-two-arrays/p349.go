package main

func intersection(nums1 []int, nums2 []int) []int {
	h1 := make(map[int]int, len(nums1))
	for _, num := range nums1 {
		h1[num] = 1
	}
	ret := make([]int, 0)
	for _, num := range nums2 {
		if h1[num] == 1 {
			h1[num] = 2
			ret = append(ret, num)
		}
	}
	return ret

}
