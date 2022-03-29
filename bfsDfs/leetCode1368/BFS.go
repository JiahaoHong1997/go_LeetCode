package leetcode1368

import "math"

func minCostBFS(grid [][]int) int {

	m, n := len(grid), len(grid[0])
	c := make([][]int,m)
	for i:=0; i<m; i++ {
		c[i] = make([]int,n)
	}
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			c[i][j] = math.MaxInt32
		}
	}
	c[0][0] = 0

	dh := []int{0, 0, 1, -1}
	dw := []int{1, -1, 0, 0}
	q := make([]int,1)
	q[0] = 0

	for len(q) > 0 {
		p := make([]int,0)
		for i:=0; i<len(q); i++ {
			h, w := q[i]/n, q[i]%n
			for k:=1; k<=4; k++ {
				mh, mw := h+dh[k-1], w+dw[k-1]
				if mh >= 0 && mh < m && mw >= 0 && mw < n {
					var x int
					if k == grid[h][w] {
						x = c[h][w]
					} else {
						x = c[h][w]+1
					}
					if x < c[mh][mw] {
						c[mh][mw] = x
						p = append(p, mh*n+mw)  // 允许多次扩展同一个点。只要当前边能够更新节点的权值，就将节点再次入队。
					}
				}
			}
		}
		q = p
	}	

	return c[m-1][n-1]
}