package ci25

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	l3 := &ListNode{}
	p := l3
	for l1 != nil && l2 != nil {
		var x *ListNode
		if l1.Val < l2.Val {
			x = l1
			l1 = l1.Next
		} else {
			x = l2
			l2 = l2.Next
		}
		x.Next = nil 
		p.Next = x
		p = p.Next
	}

	if l1 != nil {
		p.Next = l1
	} else if l2 != nil {
		p.Next = l2
	}
	return l3.Next
}