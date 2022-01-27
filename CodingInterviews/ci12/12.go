package ci12

// DFS
func exist(board [][]byte, word string) bool {

    h_idx := []int{0,1,-1,0}
    w_idx := []int{1,0,0,-1}
    r, c := len(board), len(board[0])
    m := make(map[int]bool)
    var dfs func(int,int,int) bool 
    dfs = func(i, j, index int) bool {

        if board[i][j] != word[index] {
            return false
        }
        if index == len(word)-1 {
            if board[i][j] == word[index] {
                return true
            }
        }
        m[i*c+j] = true
        defer func() {
            m[i*c+j] = false
        }()

        for k:=0; k<4; k++ {
            h, w := i+h_idx[k], j+w_idx[k]
            if h >= 0 && h < r && w >= 0 && w < c && !m[h*c+w] {
                if dfs(h,w,index+1) {
                    return true
                }
            }
        }
        return false
    }  

    for i:=0; i<r; i++ {
        for j:=0; j<c; j++ {
            if board[i][j] == word[0] {
                if dfs(i,j,0) {
                    return true
                }
            }
        }
    } 
    return false
}