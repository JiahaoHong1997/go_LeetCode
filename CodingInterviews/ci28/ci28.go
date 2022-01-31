package ci28


// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
 }

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	var isEqual func([]*TreeNode) bool 
	isEqual = func(queue []*TreeNode) bool {
		n := len(queue)
		for i:=0; i<n/2; i++ {
			if queue[i] == nil && queue[n-i-1] != nil {
				return false
			} else if queue[i] != nil && queue[n-i-1] == nil {
				return false
			} else if queue[i] == nil && queue[n-i-1] == nil {
				continue
			}

			if queue[i].Val != queue[n-i-1].Val {
				return false
			}
		}
		return true
	}

	p := []*TreeNode{}
	p = append(p, root)

	for len(p) > 0 {
		q := []*TreeNode{}
		for i:=0; i<len(p); i++ {
			x := p[i]
			if x != nil {
				q = append(q, x.Left, x.Right)
			}
		}
		if !isEqual(q) {
			return false
		}
		p = q
	}
	return true
}