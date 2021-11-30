package leetcode213

// Since House[1] and House[n] are adjacent, they cannot be robbed together. 
// Therefore, the problem becomes to rob either House[1]～House[n-1] or House[2]～House[n], depending on which choice offers more money. 
// Now the problem has degenerated to the House Robber, which is already been solved.
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	} else if len(nums) == 3 {
		return max(max(nums[0], nums[1]), nums[2])
	}

	dp1 := make([]int, len(nums))
	dp2 := make([]int, len(nums))

	dp1[0], dp2[0] = nums[0], 0
	dp1[1], dp2[1] = max(nums[0], nums[1]), nums[1]
	dp2[2] = max(nums[1], nums[2])
	for i := 2; i < len(nums); i++ {
		if i < len(nums)-1 {
			dp1[i] = max(dp1[i-1], dp1[i-2]+nums[i])
		}
		if i > 2 {
			dp2[i] = max(dp2[i-1], dp2[i-2]+nums[i])
		}
	}

	return max(dp1[len(nums)-2], dp2[len(nums)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
