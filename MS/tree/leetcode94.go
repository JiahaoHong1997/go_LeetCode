/*
 * @Descripttion: 二叉树的中序遍历
 * @Solution:
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-08-26 13:49:00
 * @LastEditors: 洪笳淏
 */
package tree

func inorderTraversal(root *TreeNode) []int {
    var res []int

    var backTracking func(*TreeNode) 
    backTracking = func(p *TreeNode) {
        if p == nil {
            return
        }
        backTracking(p.Left)
        res = append(res, p.Val)
        backTracking(p.Right)
    }

    backTracking(root)
    return res
}