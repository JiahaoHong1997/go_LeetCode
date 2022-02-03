package ci24

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	first, second := head, head.Next
	for second != nil {
		next := second.Next
		second.Next = first
		first = second
		second = next
	}
	head.Next = nil
	return first
}