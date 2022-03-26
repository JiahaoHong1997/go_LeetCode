package houserobber1

func rob(nums []int) int {
    dp := make([]int,len(nums))
    dp[0] = nums[0]

    for i:=1; i<len(nums); i++ {
        if i == 1 {
            dp[i] = max(nums[0], nums[1])
            continue
        }

        dp[i] = max(dp[i-1],dp[i-2]+nums[i])
    }
    return dp[len(dp)-1]
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}