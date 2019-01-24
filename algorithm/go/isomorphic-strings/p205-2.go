package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := make(map[byte]byte, 26)
	m2 := make(map[byte]byte, 26)
	var v1, v2 byte
	var ok1, ok2 bool
	for i := 0; i < len(s); i++ {
		v1, ok1 = m1[s[i]]
		if ok1 {
			if v1 != t[i] {
				return false
			}
		} else {
			v2, ok2 = m2[t[i]]
			if ok2 && v2 != s[i] {
				return false
			}
			m1[s[i]] = t[i]
			m2[t[i]] = s[i]
		}
	}
	return true
}
