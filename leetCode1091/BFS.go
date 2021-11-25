package leetCode1091

// 根据测试的结论，使用map的时间代价要大很多
 
// 将路过的0置为1，不使用map
func shortestPathBinaryMatrix(grid [][]int) int {
    if grid[0][0] == 1 {
        return -1
    }

    if len(grid) == 1 {
        return 1
    }
    
    maxNum := BFS(grid,0,0)
    return maxNum
}

func BFS(grid [][]int, h,w int) int {
    if h == len(grid)-1 && w == len(grid[0])-1 {
        return 1
    }
    h_idx := []int{1,0,0,-1,1,1,-1,-1}
    w_idx := []int{0,1,-1,0,1,-1,1,-1}
    n := len(grid)

    hh := []int{h}
    ww := []int{w}
    maxNum := 1
    for len(hh) > 0 {
        hhh := []int{}
        www := []int{}

        for i:=0; i<len(hh); i++ {
            H,W := hh[i], ww[i]
            for i:=0; i<8; i++ {
                mh, mw := H+h_idx[i], W+w_idx[i]
                if mh >= 0 && mh < n && mw >= 0 && mw < n && grid[mh][mw] == 0 {
                    if mh == n-1 && mw == n-1 {
                        return maxNum+1
                    }
                    hhh = append(hhh,mh)
                    www = append(www,mw)
                    grid[mh][mw] = 1
                }
            }
        }
        hh, ww = hhh, www
        maxNum++
    }
    
    return -1
}

// 使用map
func shortestPathBinaryMatrixMap(grid [][]int) int {
    if grid[0][0] == 1 {
        return -1
    }

    if len(grid) == 1 {
        return 1
    }
    
    m := make(map[int]bool,0)
    m[0] = true
    maxNum := BFSMap(grid,0,0,m)

    return maxNum
}

func BFSMap(grid [][]int, h,w int, m map[int]bool) int {
    if h == len(grid)-1 && w == len(grid[0])-1 {
        return 1
    }
    h_idx := []int{1,0,0,-1,1,1,-1,-1}
    w_idx := []int{0,1,-1,0,1,-1,1,-1}
    n := len(grid)

    hh := []int{h}
    ww := []int{w}
    maxNum := 1
    for len(hh) > 0 {
        hhh := []int{}
        www := []int{}

        for i:=0; i<len(hh); i++ {
            H,W := hh[i], ww[i]
            for i:=0; i<8; i++ {
                mh, mw := H+h_idx[i], W+w_idx[i]
                if !m[mh*n+mw] && mh >= 0 && mh < n && mw >= 0 && mw < n && grid[mh][mw] == 0 {
                    if mh == n-1 && mw == n-1 {
                        return maxNum+1
                    }
                    hhh = append(hhh,mh)
                    www = append(www,mw)
                    m[mh*n+mw] = true
                }
            }
        }
        hh, ww = hhh, www
        maxNum++
    }
    
    return -1
}