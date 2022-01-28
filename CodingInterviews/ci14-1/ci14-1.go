package ci14

func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 1, 1

    for i:=3; i<=n; i++ {
        dp[i] = dp[i-1]

        for j:=1; j<=i/2; j++ {
            t1 := max(dp[j],j)
            t2 := max(dp[i-j],i-j)
            dp[i] = max(dp[i], t1*t2)
        }

    }
    return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}