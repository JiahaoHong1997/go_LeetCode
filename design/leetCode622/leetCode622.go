package leetcode622

type MyCircularQueue struct {
	length 	int
	front 	int
	rear 	int
	full 	bool
	arr 	[]int
}


func Constructor(k int) MyCircularQueue {
	a := make([]int, k)
	return MyCircularQueue{
		length: k,
		front: 0,
		rear: 0,
		full: false,
		arr: a,
	}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.full {
		return false
	}

	this.arr[this.rear] = value
	this.rear = (this.rear + 1) % this.length
	if this.rear == this.front {
		this.full = true
	}
	return true
}


func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.front = (this.front + 1) % this.length
	if this.full {
		this.full = false
	}
	return true
}


func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.arr[this.front]
}


func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	var t int
    if this.rear == 0 {
        t = this.length - 1
    } else {
        t = this.rear - 1
    }
	return this.arr[t]
}


func (this *MyCircularQueue) IsEmpty() bool {
	return !this.full && this.front == this.rear
}


func (this *MyCircularQueue) IsFull() bool {
	return this.full
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */