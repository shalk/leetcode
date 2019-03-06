package main

func reverseVowels(s string) string {
	if s == "" {
		return ""
	}
	ret := []byte(s)
	for i, j := 0, len(s)-1; i < j; {
		if !isVowel(ret[i]) {
			i++
		} else if !isVowel(ret[j]) {
			j--
		} else {
			ret[i], ret[j] = ret[j], ret[i]
			i++
			j--
		}
	}
	return string(ret)
}

func isVowel(s byte) bool {
	var arr = []byte("aeiouAEIOU")
	for _, v := range arr {
		if s == v {
			return true
		}
	}
	return false
}
