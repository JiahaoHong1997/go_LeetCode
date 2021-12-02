package leetcode413

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	n := len(nums)
	dp := make([][]bool, n)
	for i:=0; i<n; i++ {
		dp[i] = make([]bool, n)
	}

	ret := 0
	for L:=3; L<=n; L++ {
		for i:=0; i<=n-L; i++ {
			j := i+L-1

			if L == 3 {
				if nums[i]-nums[i+1] == nums[i+1]-nums[i+2] {
					dp[i][j] = true
				}
			} else if L == 4 {
				if nums[i]-nums[i+1] == nums[j-1]-nums[j] {
					dp[i][j] = nums[i]-nums[i+1] == nums[i+1]-nums[i+2]
				}
			} else {
				if nums[i]-nums[i+1] == nums[j-1]-nums[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] {
				ret++
			}
		}
		
	}
	return ret
}