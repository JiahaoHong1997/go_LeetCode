/*
 * @Descripttion: 最大矩形
给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

示例 1：
输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：6
解释：最大矩形如上图所示。

示例 2：
输入：matrix = []
输出：0

示例 3：
输入：matrix = [["0"]]
输出：0

示例 4：
输入：matrix = [["1"]]
输出：1

示例 5：
输入：matrix = [["0","0"]]
输出：0

提示：
rows == matrix.length
cols == matrix[0].length
1 <= row, cols <= 200
matrix[i][j] 为 '0' 或 '1'
 * @Solution: 本题是leetcode84 的升级版，总体思路就是将二维矩阵降维，再利用单调栈逐行求解
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-09-09 14:56:51
 * @LastEditors: 洪笳淏
 */
package monotonestack

func maximalRectangle(matrix [][]byte) int {
    var res int
    m, n := len(matrix), len(matrix[0])
    barHeight := make([]int, n)
    for i:=0; i<m; i++ {
        if i == 0 {
            for j:=0; j<n; j++ {
                if matrix[i][j] == '1' {
                    barHeight[j] = 1
                }
            }
        } else {
            for j:=0; j<n; j++ {
                if matrix[i][j] == '1' {
                    barHeight[j]++
                } else {
                    barHeight[j] = 0
                }
            }
        }
        
        res = max(res, maxRectangleInBar(barHeight))
    }

    return res
}

func maxRectangleInBar(nums []int) int {
    nums = append([]int{0}, append(nums, 0)...)
    stack := []int{}
    res := 0
    // fmt.Println(nums)
    for i:=0; i<len(nums); i++ {
        if len(stack) == 0 {
            stack = append(stack, i)
        } else {
            for len(stack)>0 && nums[i] < nums[stack[len(stack)-1]] {
                index := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                res = max(res, nums[index]*(i-stack[len(stack)-1]-1))
                // fmt.Println(res,i,stack)
            }
            stack = append(stack, i)
        }
    }

    return res
}
