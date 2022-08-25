/*
 * @Descripttion:
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组 是数组中的一个连续部分。
 * @Solution: 经典基础dp题
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-08-25 12:30:20
 * @LastEditors: 洪笳淏
*/
package dp

func maxSubArray(nums []int) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	res := nums[0]
	subArray := nums[0]
	for i := 1; i < len(nums); i++ {
		subArray = max(subArray+nums[i], nums[i])
		res = max(res, subArray)
	}

	return res
}
