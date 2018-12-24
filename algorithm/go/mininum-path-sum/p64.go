package main

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dp[i][j] = grid[i][j]
			} else if i == m-1 {
				dp[i][j] = grid[i][j] + dp[i][j+1]
			} else if j == n-1 {
				dp[i][j] = grid[i][j] + dp[i+1][j]
			} else {
				if dp[i][j+1] > dp[i+1][j] {
					dp[i][j] = grid[i][j] + dp[i+1][j]
				} else {
					dp[i][j] = grid[i][j] + dp[i][j+1]
				}
			}
		}
	}
	return dp[0][0]
}
