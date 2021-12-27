package leetcode25

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	
	ticker := 0
	dummy := &ListNode{}
	last := &ListNode{}
	for head != nil {
		p := head
		count := 0
		for ; count < k - 1 && head != nil; count++ {
			head = head.Next
		} 
		if count <= k - 1 && head == nil {
			break
		}
		q := head
		head = head.Next
		q, p = reverse(p,q)
		if ticker == 0 {
			ticker = 1
			dummy = q
			p.Next = head
			last = p
		} else {
			last.Next = q
			last = p
			p.Next = head
		}	
	}
	return dummy
}

func reverse(head,tail *ListNode) (*ListNode, *ListNode) {
	if head.Next == tail {
		tail.Next = head
		return tail, head
	}

	prev, cur := head, head.Next
	for prev != tail {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	} 

	return prev, head
}