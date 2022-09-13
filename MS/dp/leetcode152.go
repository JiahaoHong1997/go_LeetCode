/*
 * @Descripttion: 乘积最大子数组
给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
测试用例的答案是一个 32-位 整数。
子数组 是数组的连续子序列。

示例 1:
输入: nums = [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。

示例 2:
输入: nums = [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 
提示:
1 <= nums.length <= 2 * 10^4
-10 <= nums[i] <= 10
nums 的任何前缀或后缀的乘积都 保证 是一个 32-位 整数
 * @Solution: 动态规划，记录当前元素与之前最大值和最小值乘积，并通过比较这两个乘积与当前元素的值计算出新的最大值和最小值，并实时更新最大子数组乘积
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-09-13 22:20:22
 * @LastEditors: 洪笳淏
 */
package dp

func maxProduct(nums []int) int {
    res := nums[0]
    maxNum := nums[0]
    minNum := nums[0]

    for i:=1; i<len(nums); i++ {
        a, b := maxNum*nums[i], minNum*nums[i]
        maxNum = max(nums[i], max(a, b))
        minNum = min(nums[i], min(a, b))
        res = max(res, maxNum)
    }

    return res
}
