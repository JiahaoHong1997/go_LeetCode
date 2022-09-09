/*
 * @Descripttion: 柱状图中最大的矩形
 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。

示例 1:
输入：heights = [2,1,5,6,2,3]
输出：10
解释：最大的矩形为图中红色区域，面积为 10

示例 2：
输入： heights = [2,4]
输出： 4
 
提示：
1 <= heights.length <=10^5
0 <= heights[i] <= 10^4

 * @Solution: 单调栈：在一维数组中对每一个数找到第一个比自己小的元素。这类“在一维数组中找第一个满足某种条件的数”的场景就是典型的单调栈应用场景。
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-09-09 14:54:58
 * @LastEditors: 洪笳淏
 */
package monotonestack

func largestRectangleArea(heights []int) int {
    stack := []int{}
    heights = append([]int{0}, append(heights, 0)...)
    var res int

    for i:=0; i<len(heights); i++ {
        if len(stack) == 0 {
            stack = append(stack, i)
        } else {
            for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
                index := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                res = max(res, heights[index]*(i-stack[len(stack)-1]-1))
            }
            stack = append(stack, i)
        }
    }

    return res
}

func max(a,b int) int {
    if a > b {
        return a
    }

    return b
}