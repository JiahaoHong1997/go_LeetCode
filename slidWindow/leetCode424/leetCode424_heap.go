package leetcode424

import "container/heap"

type intHeap []int
var a [26]int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Swap(i,j int) {
	h[i], h[j] = h[j], h[i]
}

func (h intHeap) Less(i,j int) bool {
	return a[h[i]] > a[h[j]]
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


func characterReplacementHeap(s string, k int) int {
	var (
		l = 0
		r = 0
		length = 0
		ret = 0
	)


	var h intHeap
	for i:=0; i<26; i++ {
		h = append(h, i)
	}
	heap.Init(&h)

	for r < len(s) {
		x := s[r]
		r++

        

		a[int(x)-'A']++
		heap.Fix(&h, int(x)-'A')
		length++

        // fmt.Println(a)
        // fmt.Println(length,a[h[0]])

		if length-a[h[0]] <= k {
			ret = max(ret,length)
			continue
		}

		for length-a[h[0]] > k {
			remove := s[l]
			l++
			length--

			a[int(remove)-'A']--
			heap.Fix(&h, int(remove)-'A')
		}

	}
	return ret
}
