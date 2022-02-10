package ci46

func translateNumDP(num int) int {
	if num == 0 {
		return 1
	}

	arr := []int{}
	for num > 0 {
		arr = append([]int{num%10}, arr...)
		num /= 10 
	}
	if len(arr) == 1 {
		return 1
	}

	dp := make([]int, len(arr))
	dp[0] = 1
	if arr[0]*10 + arr[1] <= 25 {
		dp[1] = 2
	} else {
		dp[1] = 1
	}

	for i:=2; i<len(dp); i++ {
		x, y := arr[i-1], arr[i]
		t := x*10 + y
		if t >= 10 && t <= 25 {
			dp[i] = dp[i-2] + dp[i-1]
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[len(dp)-1]
}