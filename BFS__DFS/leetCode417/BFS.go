package leetcode417

func pacificAtlanticBFS(heights [][]int) [][]int {
	m := len(heights)
	n := len(heights[0])
	ret := make([][]int,0)

	for i:=0;i<m; i++ {
		for j:=0; j<n; j++ {
			if bfs(heights,i,j) {
				ret = append(ret, []int{i,j})
			}
		}
	}
	return ret
}

func bfs(heights [][]int, i,j int) bool {
	h_idx := []int{1,-1,0,0}
	w_idx := []int{0,0,1,-1}
	m := len(heights)
	n := len(heights[0])
	q := make([]int,0)
	q = append(q, i*n+j)
	pacific := false
	atlantic := false

	for len(q) > 0 {
		p := make([]int,0)

		for k := 0; k<len(q); k++ {
			h, w := i+h_idx[k], j+w_idx[k]
			for l := 0; l<4; l++ {
				if h >= 0 && w >= 0 && h < n && w < m {
					if heights[h][w] <= heights[i][j] {
						p = append(p, h*n+w)
					}
					if h == 0 || w == 0 {
						pacific = true
					}
					if h == n-1 || w == m-1 {
						atlantic = true
					}
					if pacific && atlantic {
						return true
					}
				}
			}
		}
		q = p
	}
	return false
}