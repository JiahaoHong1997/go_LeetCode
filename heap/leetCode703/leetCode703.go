package leetCode703
import "container/heap"

// 本题要持续不断第输出第 k 大的数，第一反应可能是建一个大端堆。但是每次都先要把前 k-1 大的数出堆，将第 k 大的数输出后还要将前面出堆的数再入堆。
// 如果建立小端堆，保证堆中元素数量为 k，每当新加入一个数，首先和堆首的元素比较，如果比堆首大，则将堆首出堆，新元素入堆；如果比堆手小，直接输出堆首元素即可
type heapInts []int

type KthLargest struct {
    topK     int
    nums     heapInts
}

func Constructor(k int, nums []int) KthLargest {
    n := len(nums)
    h := make(heapInts,n)
    copy(h,nums)
    heap.Init(&h)
    
    for h.Len() > k {
        heap.Pop(&h)
    } 
    return KthLargest {
        topK    :   k,
        nums    :   h,
    }
}

func (h heapInts) Len() int {
    return len(h)
}

func (h heapInts) Swap(i,j int) {
    h[i],h[j] = h[j],h[i]
}

func (h heapInts) Less(i,j int) bool {
    return h[i]<h[j]
}

func (h *heapInts) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *heapInts) Pop() interface{} {
    old := *h
    n := len(old)
    *h = old[:n-1]
    return old[n-1]
}

func (kl *KthLargest) Add(val int) int {

    if len(kl.nums) < kl.topK  {
        heap.Push(&kl.nums,val)
    } else if val > kl.nums[0] {
        heap.Pop(&kl.nums)
        heap.Push(&kl.nums,val)
    }
    
    return kl.nums[0]
}