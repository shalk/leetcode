package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	m := len(board)
	if m != 9 {
		return false
	}
	if len(board[0]) != 9 {
		return false
	}
	// check row
	for i := 0; i < 9; i++ {
		used := make([]bool, 9)
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if used[board[i][j]-'1'] {
				// fmt.Printf("i: %d, j: %d \n", i, j)
				return false
			} else {
				// fmt.Printf("i: %d, j: %d \n", i, j)
				used[board[i][j]-'1'] = true
			}

		}
	}

	// check col
	for j := 0; j < 9; j++ {
		used := make([]bool, 9)
		for i := 0; i < 9; i++ {
			if board[i][j] == '.' {
				continue
			}
			if used[board[i][j]-'1'] {
				return false
			} else {
				used[board[i][j]-'1'] = true
			}

		}
	}

	// check 3x3
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			used := make([]bool, 9)
			for a := i; a < i+3; a++ {
				for b := j; b < j+3; b++ {
					if board[a][b] == '.' {
						continue
					}
					if used[board[a][b]-'1'] {
						return false
					} else {
						used[board[a][b]-'1'] = true
					}
				}
			}
		}
	}
	return true
}
