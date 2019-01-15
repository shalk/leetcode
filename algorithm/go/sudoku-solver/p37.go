package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func solveSudoku(board [][]byte) {
	row := make([][]bool, 9)
	col := make([][]bool, 9)
	box := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		row[i] = make([]bool, 9)
		col[i] = make([]bool, 9)
		box[i] = make([]bool, 9)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				row[i][board[i][j]-'1'] = true
				col[j][board[i][j]-'1'] = true
				box[(i/3)*3+j/3][board[i][j]-'1'] = true
			}
		}
	}

	helper(board, row, col, box)
}

func helper(board [][]byte, row [][]bool, col [][]bool, box [][]bool) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for k := '1'; k <= '9'; k++ {
					if row[i][k-'1'] || col[j][k-'1'] || box[(i/3)*3+j/3][k-'1'] {
						continue
					}
					row[i][k-'1'] = true
					col[j][k-'1'] = true
					box[(i/3)*3+j/3][k-'1'] = true
					board[i][j] = byte(k)

					if helper(board, row, col, box) {
						return true
					}
					row[i][k-'1'] = false
					col[j][k-'1'] = false
					box[(i/3)*3+j/3][k-'1'] = false
					board[i][j] = '.'
				}
				return false
			}
		}
	}
	return true
}
