package leetcode1368

func minCostDFS(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if m == 1 && n == 1 {
		return 0
	}
	traveled := make([][]bool, m)
	for i := 0; i < m; i++ {
		traveled[i] = make([]bool, n)
	}

	cost := make([][]int, m)
	for i := 0; i < m; i++ {
		cost[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cost[i][j] = m + n - 2  // 设置每个位置的最大消耗次数，当大于到该位置的消耗则回溯
		}
	}
	dh := []int{0, 0, 1, -1}
	dw := []int{1, -1, 0, 0}

	var DFS func(int, int, int) int
	DFS = func(c, i, j int) int {
		traveled[i][j] = true
		defer func() {
			traveled[i][j] = false
		}()
		if c >= cost[i][j] {
			return c
		}
		if i == m-1 && j == n-1 {
			return c
		}

		for k := 1; k <= 4; k++ {
			h, w := i+dh[k-1], j+dw[k-1]
			if h >= 0 && h < m && w >= 0 && w < n && !traveled[h][w] {
				var x int
				if grid[i][j] == k {
					x = DFS(c, h, w)
				} else {
					x = DFS(c+1, h, w)
				}

				if x < cost[h][w] {
					cost[h][w] = x
				}
			}
		}
		return c
	}

	DFS(0, 0, 0)
	return cost[m-1][n-1]

}
