package ci35

// Definition for a Node.
type Node struct {
    Val int
    Next *Node   
	Random *Node
}


func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)

	var backTracking func(*Node) *Node
	backTracking = func(p *Node) *Node {
		if p == nil {
			return nil
		}
		if q, ok := m[p]; ok {
			return q
		}

		newNode := &Node{Val: p.Val}
		m[p] = newNode
		newNode.Next = backTracking(p.Next)
		newNode.Random = backTracking(p.Random)		
		return newNode
	}

	return backTracking(head)
}