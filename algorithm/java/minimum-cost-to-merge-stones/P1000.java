class Solution {
    public int mergeStones(int[] stones, int K) {
        int n = stones.length;
        if ( n % (K-1) != 1 % (K-1)) {
            return -1;
        }
        int[][] dp = new int[n][n];
        int[] presum = new int[n + 1];
        for (int i = 1; i <= n; i++) {
            presum[i] = presum[i - 1] + stones[i - 1];
        }
        for (int i = n - K; i >= 0; i--) {
            for (int j = i + K - 1; j < n; j++) {
                int cost = Integer.MAX_VALUE;
                int len = j - i + 1;
                int s = len % (K - 1) == 1 % (K - 1) ? len - (K - 1) : len;
                for (int k = i + 1; k <= j; k++) {
                    int leftLen = k - i;
                    int rightLen = j - k + 1;
                    if ((leftLen - 1) / (K - 1) + (rightLen - 1) / (K - 1) == (s - 1) / (K - 1)) {
                        cost = Math.min(cost, dp[i][k - 1] + dp[k][j]);
                    }
                }
                // merge together
                if (len % (K - 1) == 1 % (K - 1)) {
                    cost += (presum[j + 1] - presum[i]);
                }
                dp[i][j] = cost;
            }
        }

        // update nums;
        return dp[0][n - 1];
    }
}
