package leetcode426

type node struct {
	Val   int
	Left  *node
	Right *node
}

func treeToDoublyList(root *node) *node {

	if root == nil {
		return root
	} else if root.Left == nil && root.Right == nil {
		root.Left = root
		root.Right = root
		return root
	}

	head := &node{Val: 1001}
	var pre, right *node

	var backTracking func(*node)
	backTracking = func(p *node) {
		if p == nil {
			return
		}

		backTracking(p.Left)

		if head.Val == 1001 {
			head = p
			pre = p
			right = p.Right
		} else if head.Val == p.Val {
			return
		} else {
			pre.Right = p
			p.Left = pre
			right = p.Right
			p.Right = head
			head.Left = p
			pre = p
		}

		backTracking(right)
	}

	backTracking(root)
	return head
}
