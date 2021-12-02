package longestPalindromeSequence

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i:=0; i<n; i++ {
		dp[i] = make([]int, n)
	}

	for i:=0; i<n; i++ {
		dp[i][i] = 1
	}

	for L:=2; L<=n; L++ {
		for i:=0; i<= n-L; i++ {
			j := i+L-1

			if L==2 {
				if s[i] == s[j] {
					dp[i][j] = 2
				} else {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1] + 2
				} else {
					dp[i][j] = max(dp[i+1][j], dp[i][j-1])
				}
			}
		}
	}
	return dp[0][n-1]
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}