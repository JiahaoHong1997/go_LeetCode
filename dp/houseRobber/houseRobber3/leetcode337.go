package houserobber3

 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
 
 func rob(root *TreeNode) int {

    choosed := make(map[*TreeNode]int)
    noChoosed := make(map[*TreeNode]int)

    var backTracking func(*TreeNode)
    backTracking = func(p *TreeNode) {
        if p == nil {
            return 
        }

        backTracking(p.Left)
        backTracking(p.Right)
        choosed[p] = noChoosed[p.Left]+noChoosed[p.Right]+p.Val
        noChoosed[p] = max(choosed[p.Left],noChoosed[p.Left])+max(choosed[p.Right],noChoosed[p.Right])
    }

    backTracking(root)
    return max(choosed[root],noChoosed[root])
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}