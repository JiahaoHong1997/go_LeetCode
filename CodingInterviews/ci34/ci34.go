package ci34

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {

	path := []int{}
	ret := [][]int{}

	var helper func(*TreeNode, int)
	helper = func(p *TreeNode, targetSum int) {

        if p == nil {
            return
        }
		if p.Left == nil && p.Right == nil {
			if targetSum == p.Val {
				t := make([]int, len(path))
				copy(t, path)
                t = append(t, p.Val)
				ret = append(ret, t)
			} 
			return
		}

        path = append(path, p.Val)
		helper(p.Left, targetSum-p.Val)
		helper(p.Right, targetSum-p.Val)
		path = path[:len(path)-1]
	}

	helper(root, target)
	return ret
}