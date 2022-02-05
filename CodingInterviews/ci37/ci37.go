package ci37

import "strconv"

type TreeNode struct {
	Val int
    Left *TreeNode
    Right *TreeNode
}

type Codec struct {
    s		string
	null 	string
	sep		string
}

func Constructor() Codec {
	return Codec{
		null: "#",
		sep: ",",
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
		this.s = this.s + this.null + this.sep
		return this.s
	}

	q := []*TreeNode{}
	q = append(q, root)
	nodeNum := 1 
	for nodeNum > 0 {
		p := []*TreeNode{}
		for i:=0; i<len(q); i++ {
			x := q[i]

			if x == nil {
				this.s = this.s + this.null + this.sep
				continue
			}
			
			this.s = this.s + strconv.Itoa(x.Val) + this.sep
			if x.Left == nil {
				p = append(p, nil)
			} else {
				nodeNum++
				p = append(p, x.Left)
			}

			if x.Right == nil {
				p = append(p, nil)
			} else {
				nodeNum++
				p = append(p, x.Right)
			}
		}
		q = p
	}
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    
}