package leetcode417

func pacificAtlanticDFS(heights [][]int) [][]int {
	m := len(heights)
	n := len(heights[0])
	ret := make([][]int,0)

	for i:=0;i<m; i++ {
		for j:=0; j<n; j++ {
			pacific := false
			atlantic := false
			if dfs(heights,i,j,&pacific,&atlantic) {
				ret = append(ret, []int{i,j})
			}
		}
	}
	return ret
}

func dfs(height [][]int, i,j int, p,a *bool) bool {
	
	if i == 0 || j == 0 {
		*p = true
	} else if i == len(height)-1 || j == len(height[0])-1 {
		*a = true
	}

	if *p && *a {
		return true
	}

	m := len(height)
	n := len(height[0])
	h_idx := []int{1,-1,0,0}
	w_idx := []int{0,0,1,-1}

	for k:=0; k<4; k++ {
		h, w := i+h_idx[k], j+w_idx[k]
		if h > 0 && w > 0 && h < m-1 && w < n-1 && height[h][w] <= height[i][j] {
			dfs(height, h, w, p, a)
		}
	}
	return false
}