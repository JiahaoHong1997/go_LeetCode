package leetcode22

func generateParenthesisBFS(n int) []string {
    return BFS(n)
}

func BFS(n int) (path []string) {
    lcount := []int{1}
    rcount := []int{0}
    path = []string{"("}

    for i:=0; i<2*n; i++ {
        l1 := []int{}
        r1 := []int{}
        newPath := []string{}

        for j:=0; j<len(path); j++ {
            sub := []byte(path[j])
            
            if lcount[j] < n {
                newPath = append(newPath, string(append(sub,'(')))
                l1 = append(l1,lcount[j]+1)
                r1 = append(r1,rcount[j])
            }
            if rcount[j] < lcount[j] {
                newPath = append(newPath, string(append(sub,')')))
                l1 = append(l1,lcount[j])
                r1 = append(r1,rcount[j]+1)
            }
        }
        if len(newPath) != 0 {
            path = newPath
        }
        
        lcount = l1
        rcount = r1
    }
    return
}