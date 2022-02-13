package ci54

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	
	nums := []int{}
	var helper func(p *TreeNode)
	helper = func(p *TreeNode) {

		if p == nil {
			return
		}

		helper(p.Left)
		nums = append(nums, p.Val)
		helper(p.Right)
	}
	helper(root)
	return nums[len(nums)-k]
}