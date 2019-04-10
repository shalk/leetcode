// refer https://leetcode.com/problems/regular-expression-matching/discuss/5651/Easy-DP-Java-Solution-with-detailed-Explanation
class Solution {
    public boolean isMatch(String s, String p) {
        //dp[i][j]
        int m = s.length();
        int n = p.length();
        boolean[][] dp = new boolean[m+1][n+1];
        // if s[i-1] == p[j-1], dp[i][j] == dp[i-1][j-1]
        // if s[i-1] != p[j-1] 
        //       if p[j-1] is a-z  dp[i][j] = false;
        //       if p[j-1] is .   dp[i][j] = dp[i-1][j-1]
        //       if p[j-1] is *                      
        //                 1. p[j-2] == s[i-1]  || p[j-2] == . // match zero, one, or more s[i-1];
        //                             dp[i][j] = dp[i][j-2]       zero
        //                             dp[i][j] = dp[i-1][j-2]    one
        //                             dp[i][j] = dp[i-1][j]      more
        //                 2. p[j-2] != s[i-1] && p[j-2] != .  // only be zero
        //                    dp[i][j] = dp[i][j-2]
        
        
        // dp[0][0] = true
        // when j = 0; dp[i][0] = false
        // when i = 0; dp[0][j] when dp[0][j-2] is true and  p[j-1] is *
        dp[0][0] = true;
        for (int j = 2; j < n+1; j+=2) {
            if (p.charAt(j-1) == '*' ) {
                dp[0][j] = dp[0][j-2];
            }
        }
        for (int i = 1; i < m+1; i++) {
            for (int j = 1; j < n+1; j++) {
                if (s.charAt(i-1) == p.charAt(j-1) || p.charAt(j-1) == '.') {
                    dp[i][j] = dp[i-1][j-1];
                } else if (p.charAt(j-1) == '*') {
                    if (p.charAt(j-2) == s.charAt(i-1) || p.charAt(j-2) == '.' ) {
                        dp[i][j] = dp[i][j-2] || dp[i-1][j-2] || dp[i-1][j];
                    } else {
                        dp[i][j] = dp[i][j-2];
                    }
                } else {
                    
                }
            }
        }
        return dp[m][n];
    }
}
