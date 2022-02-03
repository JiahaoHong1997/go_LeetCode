package ci30


type MinStack struct {
    minNum  []int
    n       int
    arr     []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack {
        minNum : make([]int,0),
        n : 0,
        arr : make([]int,0),
    }
}


func (this *MinStack) Push(x int)  {
    if this.n == 0 {
        this.minNum = append(this.minNum, x)
    } else {
        a := min(this.minNum[this.n-1], x)
        this.minNum = append(this.minNum, a)
    }
    this.arr = append(this.arr, x)
    this.n++
}


func (this *MinStack) Pop()  {
    this.arr = this.arr[:this.n-1]
    this.minNum = this.minNum[:this.n-1]
    this.n--
}


func (this *MinStack) Top() int {
    return this.arr[this.n-1]
}


func (this *MinStack) Min() int {
    return this.minNum[this.n-1]
}

func min(a, b int) int {
    if a < b {
        return a 
    }
    return b
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */