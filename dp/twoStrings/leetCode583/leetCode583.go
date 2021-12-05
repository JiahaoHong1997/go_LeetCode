package leetcode583

// 本题是 leetcode1143 的进阶题，分别删除若干字符之后使得两个字符串相同，则剩下的字符为两个字符串的公共子序列。
// 为了使删除操作的次数最少，剩下的字符应尽可能多。当剩下的字符为两个字符串的最长公共子序列时，删除操作的次数最少。
// 因此，可以计算两个字符串的最长公共子序列的长度，然后分别计算两个字符串的长度和最长公共子序列的长度之差，即为两个字符串分别需要删除的字符数，
// 两个字符串各自需要删除的字符数之和即为最少的删除操作的总次数。

func minDistance(word1 string, word2 string) int {
	dp := make([][]int,len(word1)+1)
	for i:=0; i<len(dp); i++ {
		dp[i] = make([]int,len(word2)+1)
	}

	for i:=1; i<=len(word1); i++ {
		for j:=1; j<=len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j],dp[i][j-1])
			}
		}
	}

	return (len(word1)-dp[len(word1)][len(word2)]) + (len(word2)-dp[len(word1)][len(word2)])
}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}