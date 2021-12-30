package leetcode1469

func getLonelyNodesDFS(root *TreeNode) []int {
	if root.Left == nil && root.Right == nil {
		return []int{}
	}
	ret := []int{}

	var DFS func(*TreeNode)
	DFS = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Left != nil && r.Right == nil {
			ret = append(ret, r.Left.Val)
		} else if r.Left == nil && r.Right != nil {
			ret = append(ret, r.Right.Val)
		}
		DFS(r.Left)
		DFS(r.Right)
	} 

	DFS(root)
	return ret
}