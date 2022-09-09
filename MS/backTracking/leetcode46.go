/*
 * @Descripttion: 全排列
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

示例 2：
输入：nums = [0,1]
输出：[[0,1],[1,0]]

示例 3：
输入：nums = [1]
输出：[[1]]
 
提示：
1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同
 * @Solution: 经典回溯题
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-09-08 16:50:34
 * @LastEditors: 洪笳淏
 */
package backtracking

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	var backTracking func([]int)
	backTracking = func(path []int) {
		if len(nums) == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			x := nums[i]
			path = append(path, x)
			nums = append(nums[:i], nums[i+1:]...)
			backTracking(path)
			path = path[:len(path)-1]
			nums = append(nums[:i], append([]int{x}, nums[i:]...)...)
		}
	}

	backTracking([]int{})
	return res
}
