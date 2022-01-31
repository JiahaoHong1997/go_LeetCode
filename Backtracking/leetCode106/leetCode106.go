package leetcode106


// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

 func buildTree(inorder []int, postorder []int) *TreeNode {

    if len(postorder) == 0 {
        return nil
    }

    n := len(postorder)
    x := postorder[n-1]
    root := &TreeNode{Val:x}

    t := 0
    for i:=0; i<len(inorder); i++ {
        if inorder[i] == x {
            t = i
            break
        }
    }

    root.Left = buildTree(inorder[:t], postorder[:t])
    root.Right = buildTree(inorder[t+1:], postorder[t:len(postorder)-1])
    return root
}