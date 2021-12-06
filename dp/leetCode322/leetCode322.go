package leetcode322


func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i:=0; i<=amount; i++ {
		dp[i] = amount+1
	}

	dp[0] = 0
	for i:=1; i<=amount; i++ {
		for j:=0; j<len(coins); j++ {
			if i - coins[j] < 0 {
				continue
			}

			dp[i] = min(dp[i], dp[i-coins[j]]+1)
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a< b {
		return a
	}
	return b
}