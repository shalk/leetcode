class Solution {
    
    public int count = 0;
    
    public int uniquePaths(int m, int n) {
        if (m == 1 || n == 1) {
            return 1;
        }
        // use cache
        int[][] cache = new int[m][n];
        helper(m,n, cache);
        //recursively
        return cache[m-1][n-1];
    }
    
    public int helper(int m, int n, int[][] cache) {
        if (m == 1 || n == 1) {
            cache[m-1][n-1] = 1;
            return 1;
        }
        if (cache[m-1][n-1] == 0) {
            cache[m-1][n-1] = helper(m-1, n, cache) + helper(m, n-1, cache);
        }
        return cache[m-1][n-1];
    }
    
}
