package leetcode694

func numDistinctIslands(grid [][]int) int {

	m := make(map[string]struct{})
	h_idx := []int{0,1,-1,0}
	w_idx := []int{1,0,0,-1}

	height, weight := len(grid), len(grid[0])
	s := ""
	bi, bj := 0, 0

	var dfs func(int,int)
	dfs = func(i, j int) {
		grid[i][j] = 0
		s += fmt.Sprintf("-%d-%d", i-bi, j-bj)

		for k:=0; k<4; k++ {
			h, w := i+h_idx[k], j+w_idx[k]
			if h >= 0 && h < height && w >= 0 && w < weight && grid[h][w] == 1 {
				dfs(h,w)
			}
		}
	}

	for i:=0; i<height; i++ {
		for j:=0; j<weight; j++ {
			if grid[i][j] == 1 {
				bi, bj, s = i, j, ""
				dfs(i,j)
				m[s] = struct{}{}
			}
		}
	}

	// for k, _ := range m {
	// 		fmt.Println(k)
	// }

	return len(m)
}