package main

import "fmt"

func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	if n == 0 {
		return ret
	}

	// l is square length
	// i is square No
	// val is value
	for l, i, val := n, 0, 1; val <= n*n; i++ {
		if l == 1 {
			ret[i][i] = val
			break
		}

		for k := 0; k < l-1; k++ {
			ret[i][i+k] = val
			val++
		}
		for k := 0; k < l-1; k++ {
			ret[i+k][i+l-1] = val
			val++
		}
		for k := 0; k < l-1; k++ {
			ret[i+l-1][i+l-1-k] = val
			val++
		}
		for k := 0; k < l-1; k++ {
			ret[i+l-1-k][i] = val
			val++
		}
		l -= 2
	}

	return ret
}
func main() {
	fmt.Println("vim-go")
}
