package leetcode253

import (
	"container/heap"
	"sort"
)

type intHeap []int

func (h intHeap) Len() int {
    return  len(h)
}

func (h intHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h intHeap) Less(i, j int) bool {
    return h[i] < h[j]
}

func (h *intHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func minMeetingRooms(intervals [][]int) int {
    sort.Slice(intervals, func(i,j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] < intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })

    var h intHeap
    heap.Push(&h, intervals[0][1])
    room := 1

    for i:=1; i<len(intervals); i++ {
        if intervals[i][0] >= h[0] {
            heap.Pop(&h)
        } else {
            room++
        }
        heap.Push(&h, intervals[i][1])
    }

    return room
}