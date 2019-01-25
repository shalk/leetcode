package main

import "fmt"

func findSubstring(s string, words []string) []int {
	ret := make([]int, 0)
	if len(words) == 0 {
		return ret
	}

	// n := len(words[0])
	wordNum := len(words)
	wordLen := len(words[0])

	wordsMap := make(map[string]int)
	for _, word := range words {
		wordsMap[word]++
	}

	for offset := 0; offset < wordLen; offset++ {
		start, p := offset+0, offset+0
		m2 := make(map[string]int)
		for start+wordNum*wordLen <= len(s) && p+wordLen <= len(s) {
			// current word
			currentWord := s[p : p+wordLen]
			// if current word in words
			if count, ok := wordsMap[currentWord]; ok {
				m2[currentWord]++
				p += wordLen
				for m2[currentWord] > count {
					m2[s[start:start+wordLen]]--
					start += wordLen
				}
			} else {
				// clear m2
				m2 = make(map[string]int)
				// try next word
				start = p + wordLen
				p = p + wordLen

			}
			if p-start == wordNum*wordLen {
				ret = append(ret, start)
			}
		}
	}
	return ret
}
