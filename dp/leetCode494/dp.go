package leetcode494

func findTargetSumWaysDP(nums []int, target int) int {

	sum := 0
	for _, v := range nums {
		sum += v
	}
	diff := sum - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}

	dp := make([][]int, len(nums)+1)
	for i:=0; i<=len(nums); i++ {
		dp[i] = make([]int, diff/2+1)
	}

	dp[0][0] = 1

	for i:=0; i < len(nums); i++ {
		for j:=0; j <= diff/2; j++ {
			if j < nums[i] {
				dp[i+1][j] = dp[i][j]
			} else {
				dp[i+1][j] = dp[i][j] + dp[i][j-nums[i]]
			}
		} 
	} 
	return dp[len(nums)][diff/2]
}