package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func isHappy(n int) bool {
	if n <= 0 {
		return false
	}
	set := make(map[int]bool)
	for true {
		if n == 1 {
			return true
		}
		if _, ok := set[n]; ok {
			return false
		} else {
			set[n] = true
		}
		next := 0
		for n > 0 {
			next = next + (n%10)*(n%10)
			n = n / 10
		}
		n = next
	}
	return false
}
