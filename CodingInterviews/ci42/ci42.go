package ci42

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	ret := 0

	var max func(int,int) int 
	max = func(a,b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp[0], ret = nums[0], nums[0]
	for i:=1; i<len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ret = max(ret, dp[i])
	}

	return ret
}