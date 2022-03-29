package leetcode130

func solve(board [][]byte)  {
    m := len(board)
    n := len(board[0])
    boundary := make(map[int]bool,0)

    for i:=0; i<n; i++ {
        if board[0][i] == 'O' {
            boundary[i] = true
        }

		if board[m-1][i] == 'O' {
			boundary[(m-1)*n+i] = true
		}
    }

	for i:=1; i<m-1; i++ {
		if board[i][0] == 'O' {
			boundary[i*n] = true
		}

		if board[i][n-1] == 'O' {
			boundary[(i+1)*n-1] = true
		} 
	}

	for k,_ := range boundary {
		BFS(board, k, boundary)
	}

	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func BFS(board [][]byte, index int, boundary map[int]bool) {
	
	n := len(board[0])
	q := []int{index}
	h_idx := []int{1,0,0,-1}
	w_idx := []int{0,1,-1,0}

	for len(q) > 0 {
		p := []int{}

		for i:=0; i<len(q); i++ {
			x := q[i]
			r, c := x/n, x%n
			board[r][c] = 'A'
			
			for j:=0; j<4; j++ {
				H, W := r+h_idx[j], c+w_idx[j]
				if H >=0 && W >= 0 && H < len(board) && W < n && board[H][W] == 'O' && !boundary[H*n+W] {
					p = append(p, H*n+W)
				}
			}
		}
		q = p
	}
}

