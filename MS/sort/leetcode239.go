/*
 * @Descripttion: 滑动窗口最大值
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回 滑动窗口中的最大值 。

示例 1：
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

示例 2：
输入：nums = [1], k = 1
输出：[1]

提示：
1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
1 <= k <= nums.length
 * @Solution: 使用堆排序（大端堆），堆中元素保存数组的索引号，排序规则是按堆中索引对应的数组元素从大到小排列
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-09-08 16:33:27
 * @LastEditors: 洪笳淏
 */
package sort

import (
	"container/heap"
)

type intHeap []int
var a []int

func(h intHeap) Len() int {
    return len(h)
}

func(h intHeap) Less(i,j int) bool {
    return a[h[i]] > a[h[j]]
}

func(h intHeap) Swap(i,j int) {
    h[i], h[j] = h[j], h[i]
}

func(h *intHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func(h *intHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func maxSlidingWindow(nums []int, k int) []int {
    var res []int
    a = nums
    var h intHeap
    for i:=0; i<k; i++ {
        heap.Push(&h, i)
    }
    res = append(res, a[h[0]])

    for i:=k; i<len(a); i++ {
        heap.Push(&h, i)
        for h[0] <= i-k {
            heap.Pop(&h)
        }
        res = append(res, a[h[0]])
    }

    return res
}