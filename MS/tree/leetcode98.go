/*
 * @Descripttion: 验证二叉搜索树
 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
 有效 二叉搜索树定义如下：
1.节点的左子树只包含 小于 当前节点的数。
2.节点的右子树只包含 大于 当前节点的数。
3.所有左子树和右子树自身必须也是二叉搜索树。

 * @Solution: 如上，满足三个条件
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-08-26 13:13:35
 * @LastEditors: 洪笳淏
*/
package tree

import "math"

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil && root.Left.Val >= root.Val {
		return false
	}
	if root.Right != nil && root.Right.Val <= root.Val {
		return false
	}

	return (isValidBST(root.Left) && root.Val > maxNum(root.Left)) && (isValidBST(root.Right) && root.Val < minNum(root.Right))
}

func maxNum(p *TreeNode) int {
	if p == nil {
		return math.MinInt32 - 1
	}
	for p.Right != nil {
		p = p.Right
	}
	return p.Val
}

func minNum(p *TreeNode) int {
	if p == nil {
		return math.MaxInt32 + 1
	}
	for p.Left != nil {
		p = p.Left
	}
	return p.Val
}
