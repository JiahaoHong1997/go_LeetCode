package weekCamp

import "sort"

func maximumScore(scores []int, edges [][]int) int {

	type node struct {
		to int
		score int
	}

	e := make([][]node, len(scores))
	for i:=0; i<len(edges); i++ {
		x, y := edges[i][0], edges[i][1]
		e[x] = append(e[x], node{y, scores[y]})
		e[y] = append(e[y], node{x, scores[x]})
	}

	for i:=0; i<len(e); i++ {
		sort.Slice(e[i], func(k, j int) bool {
			return e[i][k].score > e[i][j].score
		})

		if len(e[i]) > 3 {
			e[i] = e[i][:3]
		}
	}
	
	ret := -1
	for i:=0; i<len(edges); i++ {
		x, y := edges[i][0], edges[i][1]
		for j:=0; j<len(e[x]); j++ {
			for k:=0; k<len(e[y]); k++ {
				if e[x][j].to != e[y][k].to && e[x][j].to != y && e[y][k].to != x {
					ret = max(ret, e[x][j].score+scores[x]+scores[y]+e[y][k].score)
				}
			}
		}
	}

	return ret
}

  

func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}
