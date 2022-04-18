package weekCamp

func minimumRounds(tasks []int) int {

	m := make(map[int]int)
	maxNum := 0

	for i:=0; i<len(tasks); i++ {
		m[tasks[i]]++
		maxNum = max(maxNum, m[tasks[i]])
	}

	if maxNum < 2 {
		return -1
	}

	dp := make([]int, maxNum+1)
	dp[0], dp[1], dp[2] = 0, 0, 1

	for i:=3; i<=maxNum; i++ {
		if i == 3 {
			dp[i] = 1
			continue
		}

		if dp[i-2] == 0 && dp[i-3] == 0 {
			dp[i] = 0
		} else if dp[i-2] == 0 {
			dp[i] = dp[i-3]+1
		} else if dp[i-3] == 0 {
			dp[i] = dp[i-2]+1
		} else {
			dp[i] = min(dp[i-2], dp[i-3])+1
		}
	}
	
	ret := 0
	for _,v := range m {
		if dp[v] == 0 {
			return -1
		}
		ret += dp[v]
	}

	return ret
}

  

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

  

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
