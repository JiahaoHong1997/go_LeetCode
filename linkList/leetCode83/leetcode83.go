package leetcode83

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	pre, cur := head, head.Next
	for cur != nil {
		next := cur.Next
		if cur.Val == pre.Val {
			pre.Next = next
			cur = next
		} else {
			pre = cur
			cur = next
		}
	}

	return head
}
