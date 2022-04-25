package weekCamp

import (
	"math"
	"sort"
)

func fullBloomFlowers(flowers [][]int, persons []int) []int {
    const MaxNum = math.MaxInt32
    pair := make([][]int, 0)
    
    for i:=0; i<len(flowers); i++ {
        begin, end := flowers[i][0], flowers[i][1]
        pair = append(pair, []int{begin, MaxNum}, []int{end, -1*MaxNum})
    }
    
    for i:=0; i<len(persons); i++ {
        pair = append(pair, []int{persons[i], i})
    }
    
    sort.Slice(pair, func(i, j int) bool {
        if pair[i][0] == pair[j][0] {
            return pair[i][1] > pair[j][1]
        }
        
        return pair[i][0] < pair[j][0]
    })
    // fmt.Println(pair)
    now := 0
    ret := make([]int, len(persons))
    for i:=0; i<len(pair); i++ {
        l := pair[i]
        if l[1] == -1*MaxNum {
            now--
        } else if l[1] == MaxNum {
            now++
        } else {
            ret[l[1]] = now
        }
    }
    
    return ret
}