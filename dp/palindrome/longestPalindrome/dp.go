package longestpalindrome

func longestpalindromeDP(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i:=0; i<n; i++ {
		dp[i] = make([]bool, n)
	}

	for i:=0; i<n; i++ {
		dp[i][i] = true
	}

	begin, last := 0, 1

	for L:=2; L<=n; L++ {
		for i:=0; i<=n-L; i++ {
			j := i+L-1

			if L == 2 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			}

			if dp[i][j] && j-i+1 > last-begin {
				begin, last = i, j+1
			}
		}
	}
	return s[begin:last]
}