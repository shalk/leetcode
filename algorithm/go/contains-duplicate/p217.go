package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		} else {
			m[num] = true
		}
	}
	return false
}
