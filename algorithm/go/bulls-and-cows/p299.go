package main

import "fmt"

func getHint(secret string, guess string) string {
	var m1 [10]int
	var m2 [10]int
	bulls := 0
	cows := 0
	for _, v := range secret {
		m1[v-'0']++
	}
	for _, v := range guess {
		m2[v-'0']++
	}

	for i := 0; i < len(secret) && i < len(guess); i++ {
		if secret[i] == guess[i] {
			bulls++
			m1[secret[i]-'0']--
			m2[secret[i]-'0']--
		}
	}
	for k, v := range m1 {
		if v != 0 {
			if m2[k] < v {
				cows += m2[k]
			} else {
				cows += v
			}
		}
	}
	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}
