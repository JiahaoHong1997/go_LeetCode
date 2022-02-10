package ci52

type ListNode struct {
    Val int
    Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	   p, q := headA, headB

	   for q != p {
			if p == nil {
				p = headB
			} else {
				p = p.Next
			}
			if q == nil {
				q = headA
			} else {
				q = q.Next
			}

	   }

	   return q
}