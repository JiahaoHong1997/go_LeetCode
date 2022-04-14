package leetcode366

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func findLeaves(root *treeNode) [][]int {

	ret := [][]int{}
	if root == nil {
		return ret
	}

	var backTracking func(*treeNode) int
	backTracking = func(p *treeNode) int {
		if p == nil {
			return -1
		}

		height := 1 + max(backTracking(p.Left), backTracking(p.Right))
		if height < len(ret) {
			ret[height] = append(ret[height], p.Val)
		} else {
			ret = append(ret, []int{p.Val})
		}

		return height
	}

	backTracking(root)

	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
