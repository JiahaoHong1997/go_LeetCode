package leetcode72

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
	}

	for i:=0; i<=len(word1); i++ {
		dp[i][0] = i
	}

	for i:=0; i<=len(word2); i++ {
		dp[0][i] = i
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1])
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

func min(a, b, c int) int {
	if a <= b && a <= c{
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}
