package leetcode53

func maxSubArray(nums []int) int {

	dp := make([]int, len(nums))
	ret := nums[0]
	dp[0] = nums[0]

	for i:=1; i<len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		ret = max(ret, dp[i])
	}
	return ret
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}