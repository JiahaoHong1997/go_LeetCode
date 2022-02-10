package ci47


func maxValueBFS(grid [][]int) int {
	h_idx := []int{0,1}
	w_idx := []int{1,0}

	m, n := len(grid), len(grid[0])
	q := []int{0}
	s := []int{grid[0][0]}
	score := make([]int, m*n)
	score[0] = grid[0][0]

	for len(q) > 0 {
		p := make([]int, 0)
		s1 := make([]int, 0)
		for i:=0; i<len(q); i++ {
			h, w := q[i]/n, q[i]%n
			ss := s[i]
			for j:=0; j<2; j++ {
				mh, mw := h+h_idx[j], w+w_idx[j]
				if mh >= 0 && mw >= 0 && mh < m && mw < n {
					p = append(p, mh*n+mw)
					s1 = append(s1, ss+grid[mh][mw])
					score[mh*n+mw] = max(score[mh*n+mw], ss+grid[mh][mw])
				}
			}
		}
		q = p
		s = s1
	}
	return score[m*n-1]
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}