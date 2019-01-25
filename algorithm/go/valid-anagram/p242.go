package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// m1 := make(map[byte]int, 26)
	// m2 := make(map[byte]int, 26)
	var m [256]int
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		m[t[i]]--
		if m[t[i]] < 0 {
			return false
		}
	}
	return true
}
