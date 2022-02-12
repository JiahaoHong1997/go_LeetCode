package ci59

import "container/heap"

type intHeap []int
var a []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i,j int) bool {
	return a[h[i]] > a[h[j]]
}

func (h intHeap) Swap(i,j int) {
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
    if len(nums) == 0 {
        return []int{}
    }
    if k == 1 {
        return nums
    }
	var h intHeap
	a = nums
	ret := []int{}
	for i:=0; i<k; i++ {
		heap.Push(&h, i)
	}

	ret = append(ret, a[h[0]])
	for i:=k; i<len(a); i++ {
		for h[0] <= i-k {
			heap.Pop(&h)
		} 
		heap.Push(&h, i)
		ret = append(ret, a[h[0]])
	}

	return ret
}