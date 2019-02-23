package main

func titleToNumber(s string) int {
	// n := len(s)
	ret := 0
	for _, v := range s {
		ret = ret*26 + int(v-'A'+1)
	}
	return ret
}
