package leetcode328

type ListNode struct {
    Val int
    Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	lastOdd := head
	lastEven := head.Next
	firstEven := head.Next
	dummy := &ListNode{Next: lastOdd}
	head = head.Next.Next
	count := 0
	for head != nil {
		p := head.Next
		if count % 2 == 0 {
			lastOdd.Next = head
			lastOdd = lastOdd.Next
			lastOdd.Next = firstEven
			lastEven.Next = p
		} else {
			lastEven = lastEven.Next
		}
		head = p
		count++
	}
	return dummy.Next
}