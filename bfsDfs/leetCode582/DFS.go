package leetcode582

func killProcessDFS(pid []int, ppid []int, kill int) []int {
	if len(pid) == 1 && kill == pid[0] {
		return pid
	}

	m := make(map[int][]int, 0)
	for i := 0; i < len(ppid); i++ {
		if _, ok := m[ppid[i]]; ok {
			m[ppid[i]] = append(m[ppid[i]], pid[i])
		} else {
			m[ppid[i]] = []int{pid[i]}
		}
	}

	ret := make([]int, 1)
	ret[0] = kill

	var DFS func(int)
	DFS = func(k int) {
		if _, ok := m[k]; !ok {
			return
		}

		l := m[k]
		for i := 0; i < len(l); i++ {
			x := l[i]
			ret = append(ret, x)
			DFS(x)
		}
	}

	DFS(kill)
	return ret
}
