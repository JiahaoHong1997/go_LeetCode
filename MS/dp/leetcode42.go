/*
 * @Descripttion: 接雨水
 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例 2：
输入：height = [4,2,0,3,2,5]
输出：9

提示：
n == height.length
1 <= n <= 2 * 10^4
0 <= height[i] <= 10^5
 * @Solution: 典中典
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-09-18 15:40:18
 * @LastEditors: 洪笳淏
*/
package dp

func trap(height []int) int {
	prefix := make([]int, len(height))
	suffix := make([]int, len(height))

	prefix[0] = height[0]
	for i := 1; i < len(height); i++ {
		prefix[i] = max(prefix[i-1], height[i])
	}

	suffix[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		suffix[i] = max(suffix[i+1], height[i])
	}

	res := 0
	for i := 0; i < len(height); i++ {
		res += min(prefix[i], suffix[i]) - height[i]
	}

	return res
}
