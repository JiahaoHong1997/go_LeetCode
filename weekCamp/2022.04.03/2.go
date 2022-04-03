package weekCamp

import (
	"sort"
)

func findWinners(matches [][]int) [][]int {
	m1 := make(map[int][]int, 0)
    m2 := make(map[int]bool, 0)
	
	for i:=0; i<len(matches); i++ {
		win, lose := matches[i][0], matches[i][1]
		m1[lose] = append(m1[lose], win)
        m2[win] = true
	} 
	
    
	ret := make([][]int, 2)
	for k, v := range m1 {
		if len(v) == 1 {
			ret[1] = append(ret[1], k)
        }
	}
    
    for k, _ := range m2 {
        if len(m1[k]) == 0 {
            ret[0] = append(ret[0], k)
        }
    }
    
    sort.Ints(ret[0])
    sort.Ints(ret[1])

	return ret
}