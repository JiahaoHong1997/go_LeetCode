package leetcode10

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i:=0; i<=len(s); i++ {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true
    matches := func(i, j int) bool {
        if i == 0 {
            return false
        }
        if p[j-1] == '.' {
            return true
        }
        return s[i-1] == p[j-1]
    }

    for i := 0; i <= len(s); i++ {
        for j := 1; j <= len(p); j++ {
            if p[j-1] == '*' {
                dp[i][j] = dp[i][j] || dp[i][j-2]
                if matches(i, j - 1) {
                    dp[i][j] = dp[i][j] || dp[i-1][j]
                }
            } else if matches(i, j) {
                dp[i][j] = dp[i][j] || dp[i-1][j-1]
            }
        }
    }
	return dp[len(s)][len(p)]
}


