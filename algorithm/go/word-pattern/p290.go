package main

import "fmt"

// func split(str string) []string {
//     ret := make([]string, 0)
//     if len(str) == 0 {
//         return ret
//     }
//     start := 0
//     end := 0
//     for end < len(str) {
//         if str[start] == ' ' {
//             start++
//             end = start
//         } else if str[end] != ' ' {
//             end++
//         } else {
//             ret = append(ret, str[start:end])
//             end++
//             start = end
//         }
//     }
//     if start != end {
//         ret = append(ret, str[start:end])
//     }
//     return ret
// }

func wordPattern(pattern string, str string) bool {
	strs := strings.Fields(str)
	if len(pattern) != len(strs) {
		return false
	}
	m1 := make(map[byte]string)
	m2 := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		if v, ok := m1[pattern[i]]; ok {
			if v != strs[i] {
				return false
			}
		} else {
			if _, ok1 := m2[strs[i]]; ok1 {
				return false
			}
			m1[pattern[i]] = strs[i]
			m2[strs[i]] = pattern[i]
		}
	}
	return true
}
