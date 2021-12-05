package leetcode1143

// 对于在两个字符串，或者是能拆解成两个字符串的问题当中，我们都定义一个二维dp数组。
func longestCommonSubsequence(text1 string, text2 string) int {
    dp := make([][]int, len(text1)+1)	// 分别以 i+1 和 j+1 为末尾元素的两个序列的最长公共部分 
	for i:=0; i<len(dp); i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	for i:=1; i<=len(text1); i++ {
		for j:=1; j<=len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1]+1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}