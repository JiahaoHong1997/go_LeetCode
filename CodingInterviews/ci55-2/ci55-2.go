package ci55

type TreeNode struct {
	Val int
    Left *TreeNode
    Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var helper func(p *TreeNode) int
	helper = func(p *TreeNode) int {

		if p == nil {
			return 0
		}

		if p.Left == nil && p.Right == nil {
			return 1
		}

		left := helper(p.Left)
		right := helper(p.Right)
		return max(left, right)+1 
	}

	l := helper(root.Left)
	r := helper(root.Right)
	if l - r <= -2 || l - r >= 2 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}