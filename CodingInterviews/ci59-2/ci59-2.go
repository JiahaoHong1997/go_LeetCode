package ci59

type MaxQueue struct {
	que 	[]int
	deque 	[]int
}


func Constructor() MaxQueue {
	return MaxQueue{
		que: []int{},
		deque: []int{},
	}
}


func (this *MaxQueue) Max_value() int {
	if len(this.deque) == 0 {
		return -1
	}
	return this.deque[0]
}


func (this *MaxQueue) Push_back(value int)  {
	this.que = append(this.que, value)
	for i:=len(this.deque)-1; i>=0; i-- {
		if this.deque[i] < value {
			this.deque = this.deque[:len(this.deque)-1]
		}
	}
	this.deque = append(this.deque, value)
}


func (this *MaxQueue) Pop_front() int {
	if len(this.que) == 0 {
		return -1
	}
	if this.que[0] == this.deque[0] {
		this.deque = this.deque[1:]
	}
	x := this.que[0]
	this.que = this.que[1:]
	return x
}


/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */