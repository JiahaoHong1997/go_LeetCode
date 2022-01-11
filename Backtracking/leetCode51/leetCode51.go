package leetcode51

func solveNQueens(n int) [][]string {
	grid := make([][]bool, n)
    b := make([][]byte, n)
	
	for i:=0; i<n; i++ {
		grid[i] = make([]bool, n)
        b[i] = make([]byte, n)
	}
	path := []int{}
    ret := [][]string{}

	var backTracking func(int)
	backTracking = func(p int) {
		times := make([][]int, n)
		for i:=0; i<n; i++ {
			times[i] = make([]int, n)
		}

		path = append(path, p)
        if len(path) == n {
            s := []string{}
            m := 0
			for j:=0; j<n; j++ {
				for k:=0; k<n; k++ {
                    if m == n {
                        b[j][k] = '.'
                    } else if j == path[m]/n && k == path[m]%n {
						b[j][k] = 'Q'
						m++  
					} else {
						b[j][k] = '.'
					}

				}
				s = append(s, string(b[j]))
			}
			ret = append(ret, s)
            path = path[:len(path)-1]
            return
        }

		r, c := p/n, p%n
		grid[r][c] = true
		times[r][c] = 1
		for i:=0; i<n; i++ {
			if i == r {
				continue
			}
			if grid[i][c] == false {
				times[i][c] = 1
				grid[i][c] = true
			}
			x := i-r
			if c-x >= 0 && c-x < n && grid[i][c-x] == false {
				times[i][c-x] = 1
				grid[i][c-x] = true
			}
		}

		for i:=0; i<n; i++ {
			if i == c {
				continue
			}
			if grid[r][i] == false {
				times[r][i] = 1
				grid[r][i] = true
			}
			x := i-c
			if r+x >= 0 && r+x < n && grid[r+x][i] == false {
				times[r+x][i] = 1
				grid[r+x][i] = true
			}
		}
		defer func() {
			for i:=0; i<n; i++ {
				for j:=0; j<n; j++ {
					if times[i][j] == 1 {
						grid[i][j] = false
						times[i][j] = 0
					}
				}
			}
            if len(path) != 0 {
                path = path[:len(path)-1]
            }
            
		}()
		
		for i:=r+1; i<n; i++ {
			for j:=0; j<n; j++ {
				if grid[i][j] == true {
					continue
				}
				backTracking(i*n+j)
			}
		}

	}

	for i:=0; i<n; i++ {
		backTracking(i)
	}
	return ret
}