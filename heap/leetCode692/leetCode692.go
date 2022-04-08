package leetcode692

import (
	"strings"
	"container/heap"
)

type wordCount struct {
	word string
	priority int
	index int
}

type priorityQueue []*wordCount

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = j
	pq[j].index = i
}

func (pq priorityQueue) Less(i, j int) bool {
	if pq[i].priority == pq[j].priority {
		if strings.Compare(pq[i].word, pq[j].word) == 1 {
			// pq[i].word 的字典序在 pq[j].word 的后面，那在优先队列中就排前面
			return true
		} else if strings.Compare(pq[i].word, pq[j].word) == -1 {
			// pq[i].word 的字典序在 pq[j].word 的前面，那在优先队列中就排后面
			return false
		}
	}
	
	return pq[i].priority < pq[j].priority
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*wordCount)
	item.index = n
	*pq = append(*pq, item)
}

  

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]

	return item
}

func topKFrequent(words []string, k int) []string {

	var pq priorityQueue
	m := make(map[string]int)
	for i:=0; i<len(words); i++ {
		m[words[i]]++
	}

	i := 0
	for w, v := range m {
		if i < k {
			wc := &wordCount {
				word : w,
				index : i,
				priority : v,
			}
			i++
			heap.Push(&pq, wc)
		} else {
			if v > pq[0].priority || (v == pq[0].priority && strings.Compare(w, pq[0].word) == -1) {
				heap.Pop(&pq)
				wc := &wordCount {
					word : w,
					index : k-1,
					priority : v,
				}
				heap.Push(&pq, wc)
			}
		}
	}

	ret := make([]string, k)
	for k > 0 {
		ret[k-1] = heap.Pop(&pq).(*wordCount).word
		k--
	}

	return ret
}