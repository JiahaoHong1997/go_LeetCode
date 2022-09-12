/*
 * @Descripttion:
 * @Solution:
 * @version:
 * @Author: 洪笳淏
 * @Date: 2022-09-12 21:37:07
 * @LastEditors: 洪笳淏
 */
/*
 * @Descripttion: 缺失的第一个正数
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

示例 1：
输入：nums = [1,2,0]
输出：3

示例 2：
输入：nums = [3,4,-1,1]
输出：2

示例 3：
输入：nums = [7,8,9,11,12]
输出：1

提示：
1 <= nums.length <= 5 * 10^5
-2^31 <= nums[i] <= 2^31 - 1
 * @Solution: 原地hash，哈希的条件如下：
 1.当前元素值不等于其索引号+1；
 2.当前元素范围在[1,len(nums)]之间，且交换后的位置的当前元素不等于当前元素
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-09-12 21:37:07
 * @LastEditors: 洪笳淏
*/
package hashmap

func firstMissingPositive(nums []int) int {

	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 && (nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i]) {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1
}
