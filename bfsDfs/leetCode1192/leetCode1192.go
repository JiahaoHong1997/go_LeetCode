package leetcode1192

func criticalConnections(n int, connections [][]int) [][]int {
	
	// 超时，哭了
    conn := make([][]int, n)
	visited := 0
	traveled := make([]bool, n)
	for i:=0; i<len(connections); i++ {
		x, y := connections[i][0], connections[i][1]
		conn[x] = append(conn[x], y)
		conn[y] = append(conn[y], x)
	}
	var DFS func(int, int, int)
	DFS = func(k,x,y int) {
		traveled[k] = true
		visited++

		l := conn[k]
		for i:=0; i<len(l); i++ {
			if (k == x && l[i] == y) || (k == y && l[i] == x) || traveled[l[i]] {
				continue
			}
			DFS(l[i],x,y)
		}
	}

	ret := [][]int{}
	for i:=0; i<len(connections); i++ {
		x, y := connections[i][0], connections[i][1]
		DFS(y,x,y)
		if visited < n {
			ret = append(ret, []int{x,y})
		}
        traveled = make([]bool, n)
		visited = 0
	}
	return ret
}