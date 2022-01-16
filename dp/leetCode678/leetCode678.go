package leetcode678

func checkValidString(s string) bool {
	dp := make([][]bool, len(s))
	for i:=0; i<len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	for i:=0; i<len(s); i++ {
		if s[i] == '*' {
			dp[i][i] = true
		}
	}

	for L := 2; L <= len(s); L++ {
		for i:=0; i <= len(s)-L; i++ {
			j := i+L-1
            if s[i] == ')' || s[j] == '(' {
				continue
			}

			if L == 2 {
				if (s[i] == '(' && s[j] == ')') || s[i] == '*' || s[j] == '*' {
					dp[i][j] = true
				}
				continue
			}

			if dp[i+1][j-1] {
				dp[i][j] = true
				continue
			}

			for k:=i; k<j; k++ {
				if dp[i][k] && dp[k+1][j] {
					dp[i][j] = true
					break
				}
			}
		}
	}
	return dp[0][len(s)-1]
}