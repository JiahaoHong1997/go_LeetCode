package ci07

type TreeNode struct {
	Val int
    Left *TreeNode
    Right *TreeNode
}
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
        return nil
    }
    x := preorder[0]
    root := &TreeNode{Val:x}

    t := 0
    for i:=0; i<len(inorder); i++ {
        if inorder[i] == x {
            t = i
            break
        }
    }

    root.Left = buildTree(preorder[1:t+1], inorder[:t])
    root.Right = buildTree(preorder[t+1:], inorder[t+1:])
    return root
}
