/*
 * @Descripttion: 搜索二维矩阵 II
 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
每行的元素从左到右升序排列。
每列的元素从上到下升序排列。

提示：
m == matrix.length
n == matrix[i].length
1 <= n, m <= 300
-10^9 <= matrix[i][j] <= 10^9
每行的所有元素从左到右升序排列
每列的所有元素从上到下升序排列
-10^9 <= target <= 10^9
 * @Solution:
 * @version:
 * @Author: 洪笳淏
 * @Date: 2022-09-12 22:13:35
 * @LastEditors: 洪笳淏
 */
package sort

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])
    left := sort.Search(m ,func(i int) bool{
        return matrix[i][0] >= target
    })
    if left == 0 && matrix[0][0] != target {
        return false
    }
    if left == m {
        left = m-1
    }

    right := sort.Search(m, func(i int) bool{
        return matrix[i][n-1] >= target
    })
    if right == m {
        return false
    }
    // fmt.Println(left,right)
    for i:=right; i<=left; i++ {
        var x int
        if x=sort.Search(n, func(j int)bool {
            return matrix[i][j] >= target
        }); x != n {
            if matrix[i][x] == target {
                return true
            }    
        }  
    }
    return false
}
