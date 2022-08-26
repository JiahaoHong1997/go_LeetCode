/*
 * @Descripttion: 全排列II
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
示例 1：
输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]

示例 2：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 
提示：
1 <= nums.length <= 8
-10 <= nums[i] <= 10
 * @Solution: 回溯
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-08-26 15:17:19
 * @LastEditors: 洪笳淏
 */
package backtracking
import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	subRes := make([]int, 0)
	var backTracking func([]int, []int)
	backTracking = func(arr, path []int) {
		if len(arr) == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(arr); i++ {
			if i > 0 && arr[i] == arr[i-1] {
				continue
			}
			x := arr[i]
			arr = append(arr[:i], arr[i+1:]...)
			path = append(path, x)
			backTracking(arr, path)
			path = path[:len(path)-1]
			arr = append(arr[:i], append([]int{x}, arr[i:]...)...)
		}
	}
	backTracking(nums, subRes)
	return res
}
