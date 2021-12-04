package leetcode300

func lengthOfLISDP(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1

	maxLength := 1
	for i:=1; i<len(nums); i++ {
		length := 1
		for j:=0; j<i; j++ {
			if nums[i] > nums[j] {
				length = max(length, dp[j]+1)
			}
		}
		dp[i] = length
		maxLength = max(maxLength, length)
	}
	return maxLength
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

