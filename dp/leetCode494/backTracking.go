package leetcode494

func findTargetSumWaysBT(nums []int, target int) int {
	ret := 0

	var backTracking func(int)
	backTracking = func(sum int) {

		if len(nums) == 0 {
			if sum == target {
				ret++
			}
			return
		}

		
		x := nums[0]
		nums = nums[1:]
		backTracking(sum+x)
		backTracking(sum-x)
		nums = append([]int{x}, nums...)
	}
	backTracking(0)
	return ret

}