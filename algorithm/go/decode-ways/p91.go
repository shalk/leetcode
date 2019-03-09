package main

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	if byte(s[n-1]) == '0' {
		dp[n-1] = 0
	} else {
		dp[n-1] = 1
	}
	dp[n] = 1
	for i := n - 2; i >= 0; i-- {
		ch0 := s[i]
		ch1 := s[i+1]
		if ch0 == '0' {
			dp[i] = 0
		} else if ch0 >= '3' {
			dp[i] = dp[i+1]
		} else if ch0 == '1' || (ch0 == '2' && ch1 <= '6') {
			dp[i] = dp[i+1] + dp[i+2]
		} else if ch0 == '2' && ch1 > '6' {
			dp[i] = dp[i+1]
		} else {

		}
	}
	return dp[0]

}
