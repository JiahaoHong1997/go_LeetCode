/*
 * @Descripttion:
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
 * @Solution: 原地旋转是难点，具体参考解题：https://leetcode.cn/problems/rotate-image/solution/xuan-zhuan-tu-xiang-by-leetcode-solution-vu3m/
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-08-25 12:22:49
 * @LastEditors: 洪笳淏
*/
package simulation

func rotate(matrix [][]int) {
	var lengthRow, lengthCloumn int
	if len(matrix)%2 == 0 {
		lengthRow, lengthCloumn = len(matrix)/2, len(matrix)/2
	} else {
		lengthRow, lengthCloumn = len(matrix)/2+1, len(matrix)/2
	}

	for i := 0; i < lengthRow; i++ {
		for j := 0; j < lengthCloumn; j++ {
			row, cloumn := i, j
			tmp := 0
			replace := matrix[row][cloumn]
			for t := 0; t < 4; t++ {
				newRow := cloumn
				newColumn := len(matrix) - row - 1
				tmp = matrix[newRow][newColumn]
				matrix[newRow][newColumn] = replace
				replace = tmp
				row, cloumn = newRow, newColumn
			}
		}
	}
}
