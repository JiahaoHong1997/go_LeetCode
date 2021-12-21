package leetcode1278

func palindromePartition(s string, k int) int {
	grid := make([][]int, len(s))
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		grid[i] = make([]int, len(s))
		dp[i] = make([]int, k)
	}

	for i := 0; i < len(s); i++ {
		grid[i][i] = 0
	}

	for j := 2; j <= len(s); j++ {
		for i := 0; i < len(s)-j; i++ {

			r := i + j - 1

			if j == 2 {
				if s[i] == s[r] {
					grid[i][r] = 0
				} else {
					grid[i][r] = 1
				}
				continue
			}
			if s[i] == s[r] {
				grid[i][r] = grid[i+1][r-1]
			} else {
				grid[i][r] = grid[i+1][r-1] + 1
			}
		}
	}

	for i := 0; i < len(s); i++ {
		dp[i][0] = grid[0][i]
		for j := 1; j < k; j++ {
			dp[i][j] = dp[i-1][j-1] + grid[i-1][i]
			for l := 2; l <= i; l++ {
				dp[i][j] = min(dp[i][j], dp[i-l][j-1]+grid[i-l][i])
			}
		}
	}
	return dp[len(s)-1][k-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
