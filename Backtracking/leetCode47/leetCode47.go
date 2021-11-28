package leetcode47

import "sort"

func permuteUnique(nums []int) [][]int {
	path := make([]int, 0)
	ret := make([][]int, 0)
	sort.Ints(nums) // 排序以方便剪枝
	backTraking(nums, path, &ret, len(nums))
	return ret
}

func backTraking(nums []int, path []int, ret *[][]int, n int) {
	if len(path) == n {
		dst := make([]int, len(path))
		copy(dst, path)
		*ret = append(*ret, dst)
		return
	}

	for i := 0; i < len(nums); i++ {

		if i >= 1 && nums[i] == nums[i-1] { // 剪枝
			continue
		}
		x := nums[i]
		path = append(path, x)
		nums = append(nums[:i], nums[i+1:]...)
		backTraking(nums, path, ret, n)
		nums = append(nums[:i], append([]int{x}, nums[i:]...)...)
		path = path[:len(path)-1]
	}
}
