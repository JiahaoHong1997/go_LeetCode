package ci68

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

   
func lowestCommonAncestorBS(root, p, q *TreeNode) *TreeNode {

	if root == nil || root == p || root == q {
		return root
	}

	if (p.Val < root.Val && root.Val < q.Val) || (p.Val > root.Val && root.Val > q.Val) {
		return root
	}

	if p.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} 
	return lowestCommonAncestor(root.Right, p, q)
}