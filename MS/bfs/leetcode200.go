/*
 * @Descripttion: 岛屿数量
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。

示例 1：
输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1

示例 2：
输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3

提示：
m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'
 * @Solution:
 * @version:
 * @Author: 洪笳淏
 * @Date: 2022-08-28 15:47:55
 * @LastEditors: 洪笳淏
 */
package bfs

func numIslands(grid [][]byte) int {

	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				findContinueIsland(grid, i, j)
			}
		}
	}

	return res
}

func findContinueIsland(grid [][]byte, h, w int) {
	h_idx := []int{0, 1, 0, -1}
	w_idx := []int{1, 0, -1, 0}
	r, c := len(grid), len(grid[0])
	q := []int{h*c + w}
	traveled := make(map[int]bool)
	traveled[h*c+w] = true
	for len(q) > 0 {
		p := make([]int, 0)
		for _, v := range q {
			ch, cw := v/c, v%c
			for d := 0; d < 4; d++ {
				mh, mw := ch+h_idx[d], cw+w_idx[d]
				if !traveled[mh*c+mw] && mh >= 0 && mh < r && mw >= 0 && mw < c && grid[mh][mw] == '1' {
					traveled[mh*c+mw] = true
					p = append(p, mh*c+mw)
					grid[mh][mw] = '0'
				}
			}
		}
		q = p
	}

	return
}
