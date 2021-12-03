package leetcode91

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 1
	dp[0][1] = 0

	for i := 1; i < len(s); i++ {
		if s[i] == '0' && (s[i-1] == '0' || int(s[i-1])-'2' > 0) { // 00 || 30
			return 0
		}

		if s[i] == '0' {	
			dp[i][0] = 0
			dp[i][1] = dp[i-1][0] 
		} else {
			if s[i-1] == '0' || int(s[i-1])-'2' > 0 || (int(s[i-1])-'2' == 0 && int(s[i])-'6' >= 1) {  // 0a || 3a || (27~29)
				dp[i][0] = dp[i-1][0] + dp[i-1][1]
				dp[i][1] = 0
			} else {
				dp[i][0] = dp[i-1][0] + dp[i-1][1]
				dp[i][1] = dp[i-1][0]
			}
		}
	}
	return dp[len(s)-1][0] + dp[len(s)-1][1]
}
