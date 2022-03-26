package houserobber2

func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    dp1 := make([]int, len(nums))
    dp2 := make([]int, len(nums))
    dp1[0], dp2[0] = nums[0], 0

    for i:=1; i<len(nums)-1; i++ {
        if i == 1 {
            dp1[i] = dp1[0]
            continue
        }
        dp1[i] = max(dp1[i-1],dp1[i-2]+nums[i])
    }

    for i:=1; i<len(nums); i++ {
        if i == 1 {
            dp2[i] = nums[1]
            continue
        }
        dp2[i] = max(dp2[i-1], dp2[i-2]+nums[i])
    }

    return max(dp1[len(nums)-2],dp2[len(nums)-1])
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}