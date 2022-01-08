package leetcode1334

import "math"

func findTheCityDijkstra(n int, edges [][]int, distanceThreshold int) int {

    // 建图
    grid := make([][]int,n)
    for i:=0; i<n; i++ {
        grid[i] = make([]int,n)
    }
    for i:=0; i<len(edges); i++ {
        l := edges[i]
        grid[l[0]][l[1]], grid[l[1]][l[0]] = l[2], l[2]
    }

    v := []int{}
    u := []int{}

    shortDis := make([][]int,n)

    var dijkstra func(int)
    dijkstra = func(k int) {
        l := make([]int,n)
        copy(l, grid[k])
        for len(u) > 0 {  
            min := math.MaxInt32
            p := 0
            index := 0
            for i:=0; i<len(u); i++ {
                x := u[i]
                if l[x] != 0 && l[x] < min {
                    min = l[x]
                    p = x
                    index = i
                }
            }

            v = append(v, p)
            u = append(u[:index], u[index+1:]...)
            for i:=0; i<len(u); i++ {
                dis := l[p]
                x := u[i]
                if l[x] == 0 && grid[p][x] != 0 {
                    l[x] = grid[p][x] + dis
                } else if l[x] != 0 && grid[p][x] != 0 && l[x] > grid[p][x] + dis {
                    l[x] = grid[p][x] + dis
                }
            }
        }
        shortDis[k] = l
    }

    for i:=0; i<n; i++ {
        v = append(v, i)
        for j:=0; j<n; j++ {
            if j != i {
                u = append(u, j)
            } 
        }
        dijkstra(i)

        v = []int{}
        u = []int{}
    } 

    ret := 0
    count := math.MaxInt32
    for i:=0; i<n; i++ {
        x := shortDis[i]
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