package leetcode55

func canJump(nums []int) bool {
	if nums[0] == 0 && len(nums) > 1 {
		return false
	}
	if nums[0] >= len(nums)-1 {
		return true
	}
	dp := make([]bool, len(nums))
	dp[0] = true
	maxJump := nums[0]
	i := 1
	for i < len(nums) {
		if i <= maxJump {
			dp[i] = true
		}
		if !dp[i] {
			return false
		}
		maxJump = max(maxJump,i+nums[i])
		i++
	}
	return true
}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}