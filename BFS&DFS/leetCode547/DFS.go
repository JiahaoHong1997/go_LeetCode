package leetcode547


func FindCircleNumDFS(isConnected [][]int) int {
    m := make(map[int]bool,0)
    count := 0

    for i:=0; i<len(isConnected); i++ {
        if m[i] {
            continue
        }
        DFS(isConnected,m,i)
        count++

    }
    return count
}

func DFS(isConnected [][]int, m map[int]bool, k int) {

    c := len(isConnected[0])
    m[k] = true

    for j:=0; j<c; j++ {
        if j == k {
            continue
        }
        if isConnected[k][j] == 1 {
            if m[j] {
                continue
            }
            DFS(isConnected,m,j)
        }
    }
}



