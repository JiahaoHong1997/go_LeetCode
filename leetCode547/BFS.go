package leetcode547

func FindCircleNumBFS(isConnected [][]int) int {
    m := make(map[int]bool,0)
    count := 0

    for i:=0; i<len(isConnected); i++ {
        if m[i] {
            continue
        }
        BFS(isConnected,i,m)
        count++
    }
    return count
}


func BFS(isConnected [][]int, k int, m map[int]bool) {
    q := []int{}
    q = append(q,k)
    m[k] = true

    for len(q) > 0 {
        p := []int{}

        for i:=0; i<len(q); i++ {
            r := q[i]
            for j:=0; j<len(isConnected[0]); j++ {
                if j != r && !m[j] && isConnected[r][j] == 1 {
                    p = append(p,j)
                    m[j] = true
                }
            }
            q = p
        }
    }
}