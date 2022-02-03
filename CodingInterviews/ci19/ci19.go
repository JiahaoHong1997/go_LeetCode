package ci19

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i:=0; i<len(dp); i++ {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true

	var match func(int, int) bool
	match = func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	for i:=0; i<=len(s); i++ {
		for j:=1; j<=len(p); j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j] || dp[i][j-2]
				if match(i, j-1) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if match(i,j) {
				dp[i][j] = dp[i][j] || dp[i-1][j-1]
			}
		}
	}
	return dp[len(s)][len(p)]
}