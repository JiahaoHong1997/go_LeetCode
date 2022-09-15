/*
 * @Descripttion: 子集
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

示例 1：
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

示例 2：
输入：nums = [0]
输出：[[],[0]]
 
提示：
1 <= nums.length <= 10
-10 <= nums[i] <= 10
nums 中的所有元素 互不相同
 * @Solution: 集合问题
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-09-15 12:55:07
 * @LastEditors: 洪笳淏
 */
package backtracking

func subsets(nums []int) [][]int {
	var res [][]int
	var backTracking func(int, []int)
	backTracking = func(start int, set []int) {

		tmp := make([]int, len(set))
		copy(tmp, set)
		res = append(res, tmp)

		for i := start; i < len(nums); i++ {
			x := nums[i]
			set = append(set, x)
			backTracking(i+1, set)
			set = set[:len(set)-1]
		}
	}

	backTracking(0, []int{})
	return res
}
