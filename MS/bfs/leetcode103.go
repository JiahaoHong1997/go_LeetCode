/*
 * @Descripttion: 二叉树的锯齿形层序遍历
 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
 * @Solution:
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-08-27 15:38:28
 * @LastEditors: 洪笳淏
*/
package bfs

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	res = append(res, []int{root.Val})
	isReverse := true
	for len(q) > 0 {
		p := make([]*TreeNode, 0)
		tmp := make([]int, 0)
		for _, v := range q {
			if v.Left != nil {
				p = append(p, v.Left)
				if isReverse {
					tmp = append([]int{v.Left.Val}, tmp...)
				} else {
					tmp = append(tmp, v.Left.Val)
				}
			}
			if v.Right != nil {
				p = append(p, v.Right)
				if isReverse {
					tmp = append([]int{v.Right.Val}, tmp...)
				} else {
					tmp = append(tmp, v.Right.Val)
				}
			}
		}
		q = p
		if isReverse {
			isReverse = false
		} else {
			isReverse = true
		}
		if len(tmp) > 0 {
			res = append(res, tmp)
		}
	}
	return res
}
