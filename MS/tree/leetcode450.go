/*
 * @Descripttion: 删除二叉搜索树中的节点
 定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。
一般来说，删除节点可分为两个步骤：
1.首先找到需要删除的节点；
2.如果找到了，删除它。
 * @Solution: 写了105行 AC，一看解题，差点晕过去
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-08-25 22:35:19
 * @LastEditors: 洪笳淏
*/
package tree

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == key {
			return nil
		} else {
			return root
		}
	}
	var findTargetNode func(*TreeNode) *TreeNode
	findTargetNode = func(p *TreeNode) *TreeNode {
		if p == nil {
			return nil
		}
		if p.Val == key {
			return p
		}

		if p.Val > key {
			return findTargetNode(p.Left)
		}
		return findTargetNode(p.Right)
	}

	target := findTargetNode(root)
	if target == nil {
		return root
	}

	var findFatherNode func(*TreeNode)
	findFatherNode = func(p *TreeNode) {
		if p == nil {
			return
		}
		if p.Left == nil && p.Right == nil {
			return
		}
		if p.Left == target {
			p.Left = nil
			return
		}
		if p.Right == target {
			p.Right = nil
			return
		}
		findFatherNode(p.Left)
		findFatherNode(p.Right)
	}

	if target.Left == nil && target.Right == nil {
		findFatherNode(root)
		return root
	}

	var findMaxNumInLeft func(*TreeNode, *TreeNode, int) int
	findMaxNumInLeft = func(p *TreeNode, preNode *TreeNode, count int) int {
		if p.Right == nil {
			x := p.Val
			if count == 1 {
				preNode.Left = preNode.Left.Left
			} else {
				preNode.Right = preNode.Right.Left
			}
			return x
		}
		return findMaxNumInLeft(p.Right, p, count+1)
	}

	var findMinNumInRight func(*TreeNode, *TreeNode, int) int
	findMinNumInRight = func(p *TreeNode, preNode *TreeNode, count int) int {
		if p.Left == nil {
			x := p.Val
			if count == 1 {
				preNode.Right = preNode.Right.Right
			} else {
				preNode.Left = preNode.Left.Right
			}

			return x
		}

		return findMinNumInRight(p.Left, p, count+1)
	}

	if target.Left == nil {
		minRight := findMinNumInRight(target.Right, target, 1)
		target.Val = minRight
		return root
	}

	maxLeft := findMaxNumInLeft(target.Left, target, 1)
	target.Val = maxLeft
	return root
}

func solutionDeleteNode(root *TreeNode, key int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case root.Val > key:
		root.Left = deleteNode(root.Left, key)
	case root.Val < key:
		root.Right = deleteNode(root.Right, key)
	case root.Left == nil || root.Right == nil:
		if root.Left == nil {
			return root.Right
		}
		return root.Left
	default:
		succ := root.Right
		for succ.Left != nil {
			succ = succ.Left
		}
		succ.Right = deleteNode(root.Right, succ.Val)
		succ.Left = root.Left
		return succ
	}
	return root
}
