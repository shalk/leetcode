class Solution {
    public boolean isMatch(String s, String p) {
        if (s == null || p == null) {
            return false;
        }
        p = trimStar(p);
        int m = s.length();
        int n = p.length();
        boolean[][] dp = new boolean[m+1][n+1];
        dp[0][0] = true;
        if (n >= 1) {
            dp[0][1] = (p.charAt(0) == '*');
        }

        for (int i = 1; i < m+1; i++) {
            for (int j = 1; j < n+1; j++) {
                char ch1 = s.charAt(i-1);
                char ch2 = p.charAt(j-1);
                if (ch2 == '?') {
                    dp[i][j] = dp[i-1][j-1];
                } else if (ch2 == '*') {
                    dp[i][j] = dp[i - 1][j] || dp[i][j - 1];            
                } else  {// ch2 is alphabet
                    if (ch2 == ch1) {
                        dp[i][j] = dp[i-1][j-1];
                    } 
                }
            }
        }
        return dp[m][n];
    }
    public String trimStar(String p) {
        StringBuffer sb = new StringBuffer();
        for (int i = 0; i < p.length(); i++) {
            char ch = p.charAt(i);
            if ( ch == '*' && i > 0 && p.charAt(i-1) == '*') {
                continue;
            } else {
                sb.append(ch);
            }
        }
        return sb.toString();
    }
}
