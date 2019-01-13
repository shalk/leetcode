package main

import "fmt"

func gameOfLife(board [][]int) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = board[i][j] * 10
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			count := 0

			if i > 0 && j > 0 && board[i-1][j-1] >= 10 {
				count++
			}

			if i > 0 && board[i-1][j] >= 10 {
				count++
			}

			if i > 0 && j < n-1 && board[i-1][j+1] >= 10 {
				count++
			}

			if j > 0 && board[i][j-1] >= 10 {
				count++
			}

			if j < n-1 && board[i][j+1] >= 10 {
				count++
			}

			if i < m-1 && j > 0 && board[i+1][j-1] >= 10 {
				count++
			}

			if i < m-1 && board[i+1][j] >= 10 {
				count++
			}

			if i < m-1 && j < n-1 && board[i+1][j+1] >= 10 {
				count++
			}
			if (board[i][j] >= 10 && (count == 2 || count == 3)) || (board[i][j] == 0 && count == 3) {
				board[i][j] += 1
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] >= 10 {
				board[i][j] -= 10
			}
		}
	}
}
