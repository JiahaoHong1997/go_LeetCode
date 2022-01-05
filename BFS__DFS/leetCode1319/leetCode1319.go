package leetcode1319


func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}

	m := make([][]int,n)
	connCount := 0
	traveled := make(map[int]bool)

	var DFS func(int)
	DFS = func(strat int) {
		traveled[strat] = true
		if len(m[strat]) == 0 {
			return
		}
		l := m[strat]
		for i:=0; i<len(l); i++ {
			x := l[i]
			if _, ok := traveled[x]; !ok {
				DFS(x)
			} 
		}
	}

	for i:=0; i<len(connections); i++ {
		x, y := connections[i][0], connections[i][1]
        m[x] = append(m[x], y)
        m[y] = append(m[y], x)
	}

	for i:=0; i<n; i++ {
		if _, ok := traveled[i]; !ok {
			DFS(i)
			connCount++
		}
	}
	return connCount-1
}