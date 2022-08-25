/*
 * @Descripttion:
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
 * @Solution: 二叉树经典题，背下来！
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-08-25 11:37:12
 * @LastEditors: 洪笳淏
*/
package tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var backTracking func(*TreeNode) *TreeNode
	backTracking = func(head *TreeNode) *TreeNode {
		if head == p || head == q {
			return head
		}
		if head == nil {
			return nil
		}

		left := backTracking(head.Left)
		right := backTracking(head.Right)
		if left != nil && right != nil {
			return head
		}
		if left == nil {
			return right
		}
		return left
	}

	return backTracking(root)
}
