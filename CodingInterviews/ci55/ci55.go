package ci55

type TreeNode struct {
	Val int
    Left *TreeNode
    Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	maxD := 0

	var max func(int, int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var helper func(p *TreeNode, h int)
	helper = func(p *TreeNode, h int) {
		
		if p == nil {
			return
		}

		if p.Left == nil && p.Right == nil {
			maxD = max(maxD, h+1)
			return
		}

		helper(p.Left, h+1)
		helper(p.Right, h+1)
	}

	helper(root, 0)
	return maxD
}