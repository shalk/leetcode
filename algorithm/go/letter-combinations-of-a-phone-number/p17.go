package main

func letterCombinations(digits string) []string {

	m := make(map[byte]string, 0)
	m['2'] = "abc"
	m['3'] = "def"
	m['4'] = "ghi"
	m['5'] = "jkl"
	m['6'] = "mno"
	m['7'] = "pqrs"
	m['8'] = "tuv"
	m['9'] = "wxyz"
	index := make([]int, len(digits))
	ret := make([]string, 0)
	if len(digits) == 0 {
		return ret
	}
	for true {
		s := make([]byte, 0)
		for i, ch := range digits {
			s = append(s, m[byte(ch)][index[i]])
		}
		ret = append(ret, string(s))
		for i := 0; i < len(index); i++ {
			// ch := digits[i]
			if index[i] < len(m[digits[i]])-1 {
				index[i]++
				break
			} else {
				index[i] = 0
			}
			if i == len(index)-1 {
				return ret
			}
		}
	}
	return ret
}
