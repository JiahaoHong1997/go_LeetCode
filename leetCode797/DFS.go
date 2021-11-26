package leetCode797

func allPathsSourceTarget(graph [][]int) [][]int {
    ret := [][]int{}
    path := []int{}

    var dfs func(int) [][]int

    dfs = func(start int) [][]int {
        if start == len(graph)-1 { // 递归出口
            path = append(path,len(graph)-1)
            dst := make([]int,len(path))
            copy(dst,path)
            ret = append(ret,dst)
        } else {
			path = append(path,start)
			// 保护现场，避免递归返回到某一层时path发生了改变
        	p := make([]int,len(path))
        	copy(p,path)

        	for i:=0; i<len(graph[start]); i++ {
            	dfs(graph[start][i])
            	path = p	// 还原现场
        	}
		}
        return ret
    }
    return dfs(0)
}