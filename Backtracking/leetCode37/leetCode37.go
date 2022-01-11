package leetcode37


func solveSudoku(board [][]byte)  {
	rowMap := make([][]bool, 9)
	cloumnMap := make([][]bool, 9)
	gridMap := make([][]bool, 9)

	for i:=0; i<9; i++ {
		rowMap[i] = make([]bool, 10)
		cloumnMap[i] = make([]bool, 10)
		gridMap[i] = make([]bool, 10)
	}
    t := 0

	var backTracking func(int,int)
	backTracking = func(m, n int) {

		for k:=1; k<=9; k++ {
			x := m/3 * 3 + n/3
			if !rowMap[m][k] && !cloumnMap[n][k] && !gridMap[x][k] {
                
				board[m][n] = byte(k + '0')
				rowMap[m][k], cloumnMap[n][k], gridMap[x][k] = true, true, true
				h, w := 100, 100
				for i:=m; i<9; i++ {
					for j:=0; j<9; j++ {
		
						if board[i][j] != '.' {
							continue
						}
						h, w = i, j
						break
					}
					if h != 100 {
						break
					}
				}
				if h != 100 {
					backTracking(h, w)
                    if t == 1 {
                        return
                    }
					rowMap[m][k], cloumnMap[n][k], gridMap[x][k] = false, false, false
					board[m][n] = '.'
				} else {
                    t = 1
					return
				}
			} 
		}
	}

	firstr, firstc := 100, 100
	for i:=0; i<9; i++ {
		for j:=0; j<9; j++ {
			if board[i][j] == '.' {
				if firstr == 100 {
					firstr, firstc = i, j
				}
				continue
			} 
			
			x := int(board[i][j]) - '0'
			rowMap[i][x] = true
			cloumnMap[j][x] = true
			y := i/3 * 3 + j / 3
			gridMap[y][x] = true
		}
	}
	backTracking(firstr, firstc)
}