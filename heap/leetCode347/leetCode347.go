package leetcode347

import "container/heap"

type intHeap []int
var m map[int]int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i,j int) bool {
	return m[h[i]] < m[h[j]]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
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

func topKFrequent(nums []int, k int) []int {

	var h intHeap
	m = make(map[int]int)
	for i:=0; i<len(nums); i++ {
		m[nums[i]]++
	}
	heap.Init(&h)

	for r, v := range m {
		if len(h) < k {
			heap.Push(&h,r)
			continue
		}
		if v > m[h[0]] {
			heap.Pop(&h)
			heap.Push(&h,r)
		}
	} 

	ret := make([]int,k)
	copy(ret, h)
	return ret
}