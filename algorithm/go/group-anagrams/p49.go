package main

import "fmt"

type sortBytes []byte

func (s sortBytes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortBytes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortBytes) Len() int {
	return len(s)
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		k := []byte(str)
		sort.Sort(sortBytes(k))
		key := string(k)

		if _, ok := m[key]; ok {
			m[key] = append(m[key], str)
		} else {
			m[key] = []string{str}
		}
	}
	ret := make([][]string, 0)
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}
