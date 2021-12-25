package leetcode239

import "container/heap"

type intHeap []int

var a []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return a[h[i]] > a[h[j]]
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

func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	var h intHeap
	i := 0
	ret := []int{}
	for ; i < k; i++ {
		h = append(h,i)
	}
    heap.Init(&h)
	ret = append(ret, a[h[0]])
	for ; i < len(nums); i++ {
		
		for len(h)>0 && i - h[0] >= k {
			heap.Pop(&h)
		} 
		heap.Push(&h, i)
		ret = append(ret, a[h[0]])
	}
	return ret
}
