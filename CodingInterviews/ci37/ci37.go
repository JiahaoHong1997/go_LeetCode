package ci37

import (
	"math"
	"strconv"
	"strings"
)

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
	return this.s
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    if len(data) == 0 {
		return nil
	}

	m := make([]int, 0)
	str := strings.Split(data, this.sep)
	for i:=0; i<len(str); i++ {
		if str[i] == this.null {
			m = append(m, math.MinInt32)
		} else {
			x, _ := strconv.Atoi(str[i])
			m = append(m, x)
		}
	}

	var helper func() *TreeNode
	helper = func() *TreeNode {
		if len(m) == 0 || m[0] == math.MaxInt32 {
			return nil
		}

		root := &TreeNode{m[0],nil,nil}
		q := []*TreeNode{root}
		for i:=1; i<len(m); {
			parent := q[0]
			q = q[1:]

			if m[i] != math.MinInt32 {
				x := &TreeNode{m[i],nil,nil}
				parent.Left = x
			} else {
				parent.Left = nil
			}

			i++

			if m[i] != math.MinInt32 {
				x := &TreeNode{m[i],nil,nil}
				parent.Right = x
			} else {
				parent.Right = nil
			}

			i++
		}
		return root
	}

	return helper()
}