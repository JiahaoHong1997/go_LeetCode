package ci47

func maxValueDFS(grid [][]int) int {

	m, n := len(grid), len(grid[0])
	ret := 0
	h_idx := []int{1,0}
	w_idx := []int{0,1}

	var backTracking func(int,int,int)
	backTracking = func(i, j, score int) {
		if i == m-1 && j == n-1 {
			ret = max(ret, score+grid[i][j])
			return
		}

		for k:=0; k<2; k++ {
			h, w := i+h_idx[k], j+w_idx[k]
			if h>=0 && w>=0 && h<m && w<n {
				backTracking(h,w,score+grid[i][j])
			}
		}
	}

	backTracking(0,0,0)
	return ret
}