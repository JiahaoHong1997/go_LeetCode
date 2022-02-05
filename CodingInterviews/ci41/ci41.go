package ci41

import "container/heap"

type intHeap1 []int
type intHeap2 []int

func (h intHeap1) Len() int {
	return len(h)
} 

func (h intHeap2) Len() int {
	return len(h)
}

func (h intHeap1) Less(i,j int) bool {
	return h[i] > h[j]
}

func (h intHeap2) Less(i,j int) bool {
	return h[i] < h[j]
}

func (h intHeap1) Swap(i,j int) {
	h[i], h[j] = h[j], h[i]
}

func (h intHeap2) Swap(i,j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap1) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap2) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap1) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *intHeap2) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}


type MedianFinder struct {
	arr1 	intHeap1
	arr2 	intHeap2
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	var h1 intHeap1
	var h2 intHeap2

	return MedianFinder{
		arr1: h1,
		arr2: h2,
	}
}


func (this *MedianFinder) AddNum(num int)  {

	if len(this.arr1) == 0 && len(this.arr2) == 0 {
		heap.Push(&this.arr1, num)
	} else if len(this.arr2) == 0 {
		var x int
		if num < this.arr1[0] {
			x = heap.Pop(&this.arr1).(int)
			heap.Push(&this.arr1, num)
			heap.Push(&this.arr2, x)
		} else {
			heap.Push(&this.arr2, num)
		}
	} else if len(this.arr1) == len(this.arr2) {
		if num > this.arr1[0] {
			heap.Push(&this.arr2, num)
		} else {
			heap.Push(&this.arr1, num)
		}
	} else if len(this.arr1) < len(this.arr2) {
		if num > this.arr2[0] {
			x := heap.Pop(&this.arr2).(int)
			heap.Push(&this.arr1, x)
			heap.Push(&this.arr2, num)
		} else {
			heap.Push(&this.arr1, num)
		}
	} else {
		if num >= this.arr1[0] {
			heap.Push(&this.arr2, num)
		} else {
			x := heap.Pop(&this.arr1).(int)
			heap.Push(&this.arr2, x)
			heap.Push(&this.arr1, num)
		}
	}
	
}

func (this *MedianFinder) FindMedian() float64 {
	var x float64
	if len(this.arr1) == len(this.arr2) {
		x = (float64(this.arr1[0]) + float64(this.arr2[0])) / 2
	} else if len(this.arr1) > len(this.arr2) {
		x = float64(this.arr1[0])
	} else {
		x = float64(this.arr2[0])
	}
	return x
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */