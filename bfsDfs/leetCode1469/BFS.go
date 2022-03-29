package leetcode1469

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getLonelyNodesBFS(root *TreeNode) []int {
	if root.Left == nil && root.Right == nil {
		return []int{}
	}
	ret := make([]int, 0)
	q := make([]*TreeNode, 1)
	q[0] = root
	for len(q) > 0 {
		p := make([]*TreeNode, 0)
		for i := 0; i < len(q); i++ {
			if q[i].Left != nil && q[i].Right != nil {
				p = append(p, q[i].Left)
				p = append(p, q[i].Right)
			} else if q[i].Left != nil {
				p = append(p, q[i].Left)
				ret = append(ret, q[i].Left.Val)
			} else if q[i].Right != nil {
				p = append(p, q[i].Right)
				ret = append(ret, q[i].Right.Val)
			}
		}
		q = p
	}
	return ret
}
