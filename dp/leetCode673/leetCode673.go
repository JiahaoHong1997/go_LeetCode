package leetcode673

// 本题是leetCode300的延伸
func findNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums))	// 索引为 i 的元素作为末尾的子序列的最大长度
	cnt := make([]int, len(nums))	// 索引为 i 的元素作为末尾的子序列

	dp[0] = 1
	cnt[0] = 1

	maxLength := 1
	ret := 1
	for i:=1; i<len(nums); i++ {
		length := 1
		cnt[i] = 1
		for j:=0; j<i; j++ {
			if nums[i] > nums[j] {
				if dp[j] + 1 == length {
					cnt[i] += cnt[j]
				} else if dp[j] + 1 > length {
					cnt[i] = cnt[j]
					length = dp[j] + 1
				}
			}
		}
		dp[i] = length
		if length > maxLength {
			ret = cnt[i]
			maxLength = length
		} else if length == maxLength {
			ret += cnt[i]
		}
	}
	return ret
}
