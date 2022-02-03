package ci22

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for k > 0 {
		fast = fast.Next
		k--
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}