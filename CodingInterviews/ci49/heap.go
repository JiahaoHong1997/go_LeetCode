package ci49

import "container/heap"

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Swap(i,j int) {
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

func nthUglyNumber(n int) int {
	var h intHeap
	heap.Push(&h, 1)
	var x int
    m := make(map[int]bool)
	for n > 0 {
		x = heap.Pop(&h).(int)
		a, b, c := 2*x, 3*x, 5*x
        if !m[a] {
            heap.Push(&h,a)
            m[a] = true
        } 
        if !m[b] {
            heap.Push(&h,b)
            m[b] = true
        }
		if !m[c] {
            heap.Push(&h,c)
            m[c] = true
        }
		
		n--
	}
	return x
}
