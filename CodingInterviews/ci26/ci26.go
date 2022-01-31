package ci26


// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	if A == nil {
        return false
    }

	var backTrcking func(p *TreeNode, q *TreeNode) bool
	backTrcking = func(p *TreeNode, q *TreeNode) bool {
		if q == nil {
			return true
		}
		if p == nil && q != nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}

		return backTrcking(p.Left, q.Left) && backTrcking(p.Right, q.Right)
	}

	if A.Val == B.Val {
		if backTrcking(A, B) {
			return true
		}
	} 
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}