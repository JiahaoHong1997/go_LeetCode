package leetcode286

// 本题使用 DFS 必超时
// goos: darwin
// goarch: amd64
// pkg: leetCode/bFS__DFS/leetCode286
// cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
// BenchmarkBFS-8               836           1430195 ns/op
// BenchmarkDFS-8                 1        132343746122 ns/op
// PASS
// ok      leetCode/bFS__DFS/leetCode286   134.440s

func wallsAndGatesDFS(rooms [][]int)  {
	doors := make([][]int, 0)
	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[0]); j++ {
			if rooms[i][j] == 0 {
				doors = append(doors, []int{i, j})
			}
		}
	}
	
	for n := 0; n < len(doors); n++ {
		dfs(&rooms, doors[n][0], doors[n][1], 1)	
        		
	}
}

func dfs(room *[][]int, h, w, step int) {
    
	m := len(*room)
	n := len((*room)[0])
	h_idx := []int{1, -1, 0, 0}
	w_idx := []int{0, 0, 1, -1}

	for i:=0; i<4; i++ {
		mh, mw := h+h_idx[i], w+w_idx[i]
		if mh >= 0 && mw >= 0 && mh < m && mw < n {
			if ((*room)[mh][mw]) != 0 && ((*room)[mh][mw]) != -1 {
                if ((*room)[mh][mw]) < step {
                    continue
                }
				(*room)[mh][mw] = step
                step++
				dfs(room,mh,mw,step)
				step--
			}
		}
	}
}