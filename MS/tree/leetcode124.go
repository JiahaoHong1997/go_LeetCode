/*
 * @Descripttion:
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给你一个二叉树的根节点 root ，返回其最大路径和。
 * @Solution:
 * @version:
 * @Author: 洪笳淏
 * @Date: 2022-08-25 23:06:59
 * @LastEditors: 洪笳淏
 */
package tree
import "math"

func maxPathSum(root *TreeNode) int {

    res := math.MinInt32
    var backTracking func(*TreeNode) int
    backTracking = func(p *TreeNode) int {
        if p == nil {
            return 0
        }

        if p.Left == nil && p.Right == nil {
            res = max(res, p.Val)
            return p.Val
        }

        left := backTracking(p.Left)
        right := backTracking(p.Right)
        now := p.Val
        if left <= 0 && right <= 0 {
            res = max(res, p.Val)
            return p.Val
        } 
        if left > 0 && right > 0 {
            res = max(res, now+left+right)
            now += max(left, right)
        } else if left > 0 {
            now += left
        } else {
            now += right
        }
        res = max(res, now)
        return now
    }
    
    backTracking(root)
    return res
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
