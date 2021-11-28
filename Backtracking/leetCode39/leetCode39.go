package leetcode39

func combinationSum(candidates []int, target int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)
	backtracking(0, candidates, target, path, &ret)
	return ret
}

func backtracking(start int, candidates []int, target int, path []int, ret *[][]int) {
	if target == 0 {
		dst := make([]int, len(path))
		copy(dst, path)
		*ret = append(*ret, dst)
		return
	} else if target < 0 {
		return
	}

	for i := start; i < len(candidates); i++ { // 从当前节点往后寻找，避免出现重复
		x := candidates[i]
		path = append(path, x)
		backtracking(i, candidates, target-x, path, ret)
		path = path[:len(path)-1]
	}
}
