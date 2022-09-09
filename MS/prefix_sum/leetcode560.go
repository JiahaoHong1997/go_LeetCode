/*
 * @Descripttion: 和为 K 的子数组
 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。

示例 1：
输入：nums = [1,1,1], k = 2
输出：2

示例 2：
输入：nums = [1,2,3], k = 3
输出：2
 
提示：
1 <= nums.length <= 2 * 10^4
-1000 <= nums[i] <= 1000
-10^7 <= k <= 10^7
 * @Solution: 使用一个map记录下不同数值的前缀和出现的次数，在遍历的过程中累积上对应的次数。注意，初始前缀和为0时次数为1
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-09-07 15:08:46
 * @LastEditors: 洪笳淏
 */
package prefixsum

func subarraySum(nums []int, k int) int {
    prefixMap := make(map[int]int)
    prefixMap[0] = 1
    var res int
    prefix := 0

    for i:=0; i<len(nums); i++ {
        prefix += nums[i]
        if v,ok := prefixMap[prefix-k]; ok {
            res += v
        }
        prefixMap[prefix] += 1
    } 
    
    return res
}