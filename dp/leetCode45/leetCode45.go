package leetcode45

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	dp := make([]int,len(nums))
	dp[0] = 0
	maxJump := nums[0]
	if maxJump >= len(nums) - 1 {
		return 1
	} 

	for i:=1; i <= maxJump; i++ {
		dp[i] = 1
	}
	
	i := 1
	for i < len(nums) {
		if i+nums[i] > maxJump {
			newMaxJump := i+nums[i]		
			if newMaxJump >= len(nums) - 1 {
				return dp[i]+1
			}
			for j:=maxJump+1; j<=newMaxJump; j++ {
				dp[j] = dp[i]+1
			}
			maxJump = newMaxJump
		}
		i++
	}

	return dp[len(nums)-1]
}

