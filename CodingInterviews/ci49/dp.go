package ci49

func nthUglyNumberDP(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	p2, p3, p5 := 0, 0, 0

	for i:=1; i<n; i++ {
		a, b, c := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(a,b),c)
		if dp[i] == a {
			p2++
		}

		if dp[i] == b {
			p3++
		}

		if dp[i] == c {
			p5++
		}
	}
	return dp[n-1]
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}