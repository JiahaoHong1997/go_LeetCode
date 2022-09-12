/*
 * @Descripttion: 在排序数组中查找元素的第一个和最后一个位置
 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
 
示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

示例 2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]

示例 3：
输入：nums = [], target = 0
输出：[-1,-1]
 
提示：
0 <= nums.length <= 10^5
-10^9 <= nums[i] <= 10^9
nums 是一个非递减数组
-10^9 <= target <= 10^9
 * @Solution: 寻找左边界和右边界
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-09-12 22:29:37
 * @LastEditors: 洪笳淏
 */
package sort

func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1,-1}
    }
    left := findLeftBoundry(nums, target)
    if left == -1 {
        return []int{-1,-1}
    }

    right := findRightBoundry(nums, target)
    return []int{left, right}
}

func findLeftBoundry(nums []int, target int) int {
    length := len(nums)
    if nums[0] > target || nums[length-1] < target {
        return -1
    }

    l, r := 0, length-1
    for l < r {
        mid := l+(r-l)>>1
        if nums[mid] == target {
            r = mid
        } else if nums[mid] > target {
            r = mid
        } else {
            l = mid+1
        }
    }
    if nums[l] != target {
        return -1
    }
    return l
}

func findRightBoundry(nums []int, target int) int {
    length := len(nums)
    if nums[0] > target || nums[length-1] < target {
        return -1
    }
    if nums[length-1] == target {
        return length-1
    }

    l, r := 0, length-1
    for l<r {
        mid := l+(r-l)>>1
        if nums[mid] == target {
            l = mid+1
        } else if nums[mid] > target {
            r = mid
        } else {
            l = mid+1
        }
    }

    if l == 0 || nums[l-1] != target {
        return -1
    }
    return l-1
}