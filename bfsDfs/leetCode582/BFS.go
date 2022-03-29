package leetcode582

func killProcessBFS(pid []int, ppid []int, kill int) []int {
	if len(pid) == 1 && kill == pid[0] {
		return pid
	}

	m := make(map[int][]int,0)
	for i:=0; i<len(ppid); i++ {
		if _, ok := m[ppid[i]]; ok {
			m[ppid[i]] = append(m[ppid[i]], pid[i])
		} else {
			m[ppid[i]] = []int{pid[i]}
		}
	}

	q := make([]int,1)
	q[0] = kill
	ret := make([]int,1)
	ret[0] = kill

	for len(q) > 0 {
		p := make([]int, 0) 
		for i:=0; i<len(q); i++ {
			x := q[i]
			ret = append(ret, m[x]...)
			p = append(p, m[x]...)
		}
		q = p
	}
	return ret
}