package leetcode1504

func numSubmat(mat [][]int) int {
    m, n := len(mat), len(mat[0])
    grid := make([][]int, m)
    for i:=0; i<m; i++ {
        grid[i] = make([]int, n) 
    }

    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            if j == 0 {
                grid[i][j] = mat[i][j]
            } else if mat[i][j] == 0  {
                continue
            } else {
                grid[i][j] = 1 + grid[i][j-1]
            }
        }
    }

    min := func(a,b int) int {
        if a < b {
            return a
        }
        return b
    }

    ret := 0
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            x := grid[i][j]
            for k:=i; k>=0 && x!=0; k-- {
                x = min(x, grid[k][j])
                ret += x
            }
        }
    }
    return ret
}