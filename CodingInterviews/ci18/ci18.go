package ci18

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}


func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}

	pre, cur := head, head.Next
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			break
		}
		pre, cur = cur, cur.Next
	}
	return head
}