/*
 * @Descripttion: 二叉树的直径
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
示例 :
给定二叉树
          1
         / \
        2   3
       / \     
      4   5    
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
注意：两结点之间的路径长度是以它们之间边的数目表示。

 * @Solution: 1.0
 * @version:
 * @Author: 洪笳淏
 * @Date: 2022-08-27 16:44:32
 * @LastEditors: 洪笳淏
 */
package tree

func diameterOfBinaryTree(root *TreeNode) int {

    var res int
    var backTracking func(*TreeNode) int
    backTracking = func(p *TreeNode) int {
        if p == nil {
            return 0
        }

        left := backTracking(p.Left)
        right := backTracking(p.Right)
        res = max(res, left+right) // 路径长度等于总节点数-1

        return max(left,right)+1
    }
    
    backTracking(root)
    return res
}