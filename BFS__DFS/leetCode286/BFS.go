package leetcode286

func wallsAndGatesBFS(rooms [][]int) {

	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[0]); j++ {
			if rooms[i][j] == 0 {
				bfs(&rooms, []int{i, j})
			}
		}
	}

}

func bfs(rooms *[][]int, target []int) {
	m := len(*rooms)
	n := len((*rooms)[0])
	h_idx := []int{1, -1, 0, 0}
	w_idx := []int{0, 0, -1, 1}

	found := make(map[int]bool)
	q := make([]int, 0)
	q = append(q, target[0]*n+target[1])
	found[target[0]*n+target[1]] = true
	step := 1

	for len(q) > 0 {
		p := make([]int, 0)
		for i := 0; i < len(q); i++ {
			h, w := q[i]/n, q[i]%n
			for j := 0; j < 4; j++ {
				mh, mw := h+h_idx[j], w+w_idx[j]
				if mh >= 0 && mh < m && mw >= 0 && mw < n && !found[mh*n+mw] {
					if (*rooms)[mh][mw] != 0 && (*rooms)[mh][mw] != -1 {
						(*rooms)[mh][mw] = min((*rooms)[mh][mw], step)
						p = append(p, mh*n+mw)
						found[mh*n+mw] = true
					}
				}
			}
		}
		step++
		q = p
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
