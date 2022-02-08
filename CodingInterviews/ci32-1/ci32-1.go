package ci32

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	
	ret := []int{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		p := []*TreeNode{}
		for i:=0; i<len(q); i++ {
			x := q[i]
			ret = append(ret, x.Val)
			if x.Left != nil {
				p = append(p, x.Left)
			}
			if x.Right != nil {
				p = append(p, x.Right)
			}
		}
		q = p
	}
	return ret
}