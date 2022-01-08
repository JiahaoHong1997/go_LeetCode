package leetcode1334

import "math"

func findTheCityFloyd(n int, edges [][]int, distanceThreshold int) int {
    // 建图
    grid := make([][]int,n)
    for i:=0; i<n; i++ {
        grid[i] = make([]int,n)
    }
    for i:=0; i<len(edges); i++ {
        l := edges[i]
        grid[l[0]][l[1]], grid[l[1]][l[0]] = l[2], l[2]
    }
    for i:=0; i<n; i++ {
        for j:=0; j<n; j++ {
            if i != j && grid[i][j] == 0 {
                grid[i][j] = math.MaxInt32
            }
        }
    }

    var floyd func()
    floyd = func() {
        for k:=0; k<n; k++ {
            for i:=0; i<n; i++ {
                for j:=0; j<n; j++ {
                    grid[i][j] = min(grid[i][j],grid[i][k]+grid[k][j])
                }
            }
        }
    }

    floyd()
    ret := 0
    count := math.MaxInt32
    for i:=0; i<n; i++ {
        x := grid[i]
        var c int
        for j:=0; j<n; j++ {
            if i == j {
                continue
            }
            if x[j] <= distanceThreshold {
                c++
            }
        }
        if c <= count {
            count = c
            ret = i
        }
    }
    return ret
}

func min(a,b int) int {
    if a<b {
        return a
    }
    return b
}