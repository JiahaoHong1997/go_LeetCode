package leetcode79

func exist(board [][]byte, word string) bool {
    h,w := len(board),len(board[0])
    m := make([][]bool,h)
    for i:=0; i<h; i++ {
        m[i] = make([]bool,w)
    }

    for i:=0; i<h; i++ {
        for j:=0; j<w; j++ {
            if backTracking(board, word, i*w+j, 0, m) {
                return true
            }
        }
    }
    return false
}

func backTracking(board [][]byte, word string, pos int, index int, m [][]bool) bool {
    n := len(board[0])
    h, w := pos/n, pos%n
    if board[h][w] != word[index] {
        return false
    }
    if index == len(word)-1 {
        return true
    }

    m[h][w] = true
    defer func() {  // 回溯到之前的节点时，需要还原被标记过的路径，将其置为未访问状态
        m[h][w] = false
    }()
    
    h_idx := []int{-1,0,0,1}
    w_idx := []int{0,-1,1,0}
    
    for j:=0; j<4; j++ {
        mh, mw := h+h_idx[j], w+w_idx[j]
        if mh >= 0 && mw >= 0 && mh < len(board) && mw < len(board[0]) && !m[mh][mw] { 
            if backTracking(board,word,mh*n+mw,index+1,m) {
                return true
            }
        }
    }
    
    return false
} 