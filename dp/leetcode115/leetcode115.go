package leetcode115

func numDistinct(s string, t string) int {
	dp := make([][]int, len(s)+1)
	for i:=0; i<len(dp); i++ {
		dp[i] = make([]int, len(t)+1)
	}

	for i:=0; i<len(s)+1; i++ {
		dp[i][len(t)] = 1
	} 

	for i:=0; i<len(t); i++ {
		dp[len(s)][i] = 0
	}

	for i:=len(s)-1; i>=0; i-- {
		for j:=len(t)-1; j>=0; j-- {
			if s[i] == t[j] {
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}

	return dp[0][0]  
}