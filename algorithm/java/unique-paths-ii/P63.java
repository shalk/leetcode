class Solution {
    public int uniquePathsWithObstacles(int[][] obstacleGrid) {
        int m = obstacleGrid.length;
        if (m == 0) {
            return 0;
        }
        int n = obstacleGrid[0].length;
        int[][] dp = new int[m][n];
        for (int i = m -1; i >= 0; i--) {
            for (int j = n - 1; j >= 0; j--) {
                if (obstacleGrid[i][j] == 1) {
                    dp[i][j] = 0;
                    continue;
                }
                if (i == m -1 && j == n-1) {            
                    dp[i][j] = 1;
                }
                // can go right
                if (j < n-1 && obstacleGrid[i][j+1] == 0 ){
                    dp[i][j] += dp[i][j+1];
                }
                // can go down
                if (i < m-1 && obstacleGrid[i+1][j] == 0) {
                    dp[i][j] += dp[i+1][j];
                }
            }
        }
        return dp[0][0];
    }
}
