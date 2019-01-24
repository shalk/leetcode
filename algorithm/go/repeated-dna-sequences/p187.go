package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return []string{}
	}
	ret := make([]string, 0, len(s)-9)
	m := make(map[string]int, len(s)-9)
	var substr string
	for i := 0; i <= len(s)-10; i++ {
		substr = s[i : i+10]
		m[substr]++
		if m[substr] == 2 {
			ret = append(ret, substr)
		}
	}
	return ret

}
