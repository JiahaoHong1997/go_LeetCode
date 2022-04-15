package leetcode474

func findMaxForm(strs []string, m int, n int) int {

	dp := make([][][]int, len(strs)+1)
	
	var numCount func(string) (int, int)
	numCount = func(s string) (int, int) {
		zeros, ones := 0, 0
		for i:=0; i<len(s); i++ {
			if s[i] == '0' {
				zeros++
			} else {
				ones++
			}
		}

		return zeros, ones
	}

	for i:=0; i<=len(strs); i++ {
		dp[i] = make([][]int, m+1)

		for j:=0; j<=m; j++ {
			dp[i][j] = make([]int, n+1)

			for k:=0; k<=n; k++ {

				if i == 0 {
					dp[i][j][k] = 0
					continue
				}

				zeros, ones := numCount(strs[i-1])
				if j<zeros || k<ones {
					dp[i][j][k] = dp[i-1][j][k]
				} else {
					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-zeros][k-ones]+1)
				}
			}
		}
	}

	return dp[len(strs)][m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}