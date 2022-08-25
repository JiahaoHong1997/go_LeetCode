/*
 * @Descripttion:
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1:
输入: [3,2,1,5,6,4], k = 2
输出: 5

示例 2:
输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4
 * @Solution: 
 方法一：快排
 方法二：堆排
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-08-25 11:18:03
 * @LastEditors: 洪笳淏
 */

package sort

func findKthLargest(nums []int, k int) int {

    var quickSort func([]int, int) int
    quickSort = func(arr []int, index int) int {
        if len(arr) <= 1 {
            return arr[0]
        }
        l, r := 0, len(arr)-1
        m := l
        for l < r {
            for arr[m] <= arr[r] && l < r {
                r--
            }
            for arr[m] >= arr[l] && l < r {
                l++
            }
            arr[l], arr[r] = arr[r], arr[l]
        }
        arr[l], arr[m] = arr[m], arr[l]
        if l == index {
            return arr[l]
        } else if l < index {
            return quickSort(arr[l+1:], index-l-1)
        } else {
            return quickSort(arr[:l], index)
        }
    }

    return quickSort(nums, len(nums)-k)
}
