package main

import "fmt"

func minWindow(s string, t string) string {
	m := make(map[byte]int)
	for _, v := range []byte(t) {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	if len(s) == 0 {
		return ""
	}

	count := len(t)
	if _, ok := m[s[0]]; ok {
		m[s[0]]--
		count--
	}
	if count == 0 {
		return string(s[0:1])
	}
	min := s
	flag := false
	for start, end := 0, 0; start < len(s) && end < len(s); {
		if count == 0 {
			flag = true
			if end-start+1 < len(min) {
				min = string(s[start : end+1])
			}
			//start++
			if _, ok := m[s[start]]; ok {
				m[s[start]]++
				if m[s[start]] > 0 {
					count++
				}
			}
			start++
		} else {
			end++
			if end >= len(s) {
				break
			}

			if _, ok := m[s[end]]; ok {
				m[s[end]]--
				if m[s[end]] >= 0 {
					count--
				}
			}
		}
	}
	if flag {
		return min
	} else {
		return ""
	}
}
