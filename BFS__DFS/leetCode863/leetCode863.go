package leetcode863

type TreeNode struct {
	Val int
	Left *TreeNode
	Right * TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	if k == 0 {
        return []int{target.Val}
    }
	ret := make([]int,0)
	m := make(map[*TreeNode]*TreeNode,0)
	traveled := make(map[*TreeNode]bool,0)

	var buildParent func(*TreeNode, *TreeNode)
	buildParent = func(p *TreeNode, c *TreeNode) {

		if c == nil {
			return
		}

		m[c] = p
		if c == root {
			buildParent(p, p.Left)
			buildParent(p, p.Right)
		} else {
			buildParent(c, c.Left)
			buildParent(c, c.Right)
		}
	}

	var findChild func(n int, p *TreeNode)
	findChild = func(n int, p *TreeNode) {

		if n == 0 && p != nil {
			ret = append(ret, p.Val)
			return
		}
		if p == nil {
			return
		}
		findChild(n-1, p.Left)
		findChild(n-1, p.Right)
	}

	var findParent func(n int, p *TreeNode)
	findParent = func(n int, p *TreeNode) {

		if n == 0 && p != nil {
			ret = append(ret, p.Val)
		}
		if p == nil {
			return
		}
		if !traveled[m[p]] {
			traveled[m[p]] = true
			findParent(n-1, m[p])
		}
		if !traveled[p.Left] {
			traveled[p.Left] = true
			findParent(n-1, p.Left)
		}
		if !traveled[p.Right] {
			traveled[p.Right] = true
			findParent(n-1, p.Right)
		}
	}
	

	p := target
	findChild(k, p)
	buildParent(root, root)
	
	traveled[target] = true
	traveled[target.Left] = true
	traveled[target.Right] = true
	findParent(k, target)



	return ret
}

