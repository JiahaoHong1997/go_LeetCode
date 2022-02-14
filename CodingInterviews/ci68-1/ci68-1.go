package ci68

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

   
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil || root == p || root == q {
		return root
	}
	
	var searchVal func(*TreeNode, int) bool 
	searchVal = func(r *TreeNode, val int) bool {

		if r == nil {
			return false
		}
		if r.Val == val {
			return true
		}

		return searchVal(r.Left, val) || searchVal(r.Right, val)
	}

	if (searchVal(root.Left, p.Val) && searchVal(root.Right, q.Val)) || (searchVal(root.Right, p.Val) && searchVal(root.Left, q.Val)) {
		return root
	} else if searchVal(root.Left, p.Val) && searchVal(root.Left, q.Val) {
		return lowestCommonAncestor(root.Left, p, q)
	} 
	return lowestCommonAncestor(root.Right, p, q)
}