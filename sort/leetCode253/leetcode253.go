package leetcode253

import (
	"math"
	"sort"
)

func minMeetingRooms(intervals [][]int) int {

    const MaxNum = math.MaxInt32
    room := 0
    ret := 0

    var max func(a, b int) int
    max = func(a, b int) int {
        if a > b {
            return a
        }

        return b
    }

    pair := make([][]int, 0)
    for i:=0; i<len(intervals); i++ {
        begin, end := intervals[i][0], intervals[i][1]
        pair = append(pair, []int{begin, MaxNum})
        pair = append(pair, []int{end, -1*MaxNum})
    }

    sort.Slice(pair, func(i, j int) bool {
        if pair[i][0] == pair[j][0] {
            return pair[i][1] < pair[j][1]
        }

        return pair[i][0] < pair[j][0]
    })

    for i:=0; i<len(pair); i++ {
        if pair[i][1] == MaxNum {
            room++
            ret = max(ret, room)
        } else if pair[i][1] == -1*MaxNum {
            room--
        }
    }

    return ret
}