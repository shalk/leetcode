package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	// move m to end;
	for i := m - 1; i >= 0; i-- {
		nums1[n+i], nums1[i] = nums1[i], nums1[n+i]
	}
	for i, j, k := n, 0, 0; i < m+n || j < n; k++ {
		if i < m+n && j < n {
			if nums1[i] < nums2[j] {
				nums1[k] = nums1[i]
				i++
			} else {
				nums1[k] = nums2[j]
				j++
			}
		} else if i < m+n {
			break
		} else {
			nums1[k] = nums2[j]
			j++
		}
	}
	return
}
