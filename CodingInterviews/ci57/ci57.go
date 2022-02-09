package ci57

func twoSum(nums []int, target int) []int {
	l, r := 0, len(nums)-1

	for l < r {
		x := nums[l] + nums[r]
		if x == target {
			return []int{nums[l],nums[r]}
		} else if x < target {
			l++
		} else {
			r--
		}
	}
	return []int{}
}