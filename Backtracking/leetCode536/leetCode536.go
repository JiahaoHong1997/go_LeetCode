package leetcode536

import "strconv"

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func str2tree(s string) *treeNode {

	if s == "" {
		return nil
	}

	var backTracking func(string) *treeNode
	backTracking = func(b string) *treeNode {

		if len(b) == 0 {
			return nil
		}

		var h *treeNode
		var index int
		var x int
		num := ""

		for ; index < len(b); index++ {
			if b[index] != '(' {
				num += b[index : index+1]
			} else {
				break
			}
		}

		x, _ = strconv.Atoi(num)
		h = &treeNode{
			Val: x,
		}

		bb := ""
		l := 0
		f := 0

		for ; index < len(b); index++ {

			if b[index] == '(' {
				if l == 0 {
					l++
					continue
				} else {
					l++
					bb += "("
				}
			} else if (b[index] >= '0' && b[index] <= '9') || b[index] == '-' {
				bb += b[index : index+1]
			} else {
				l--
				if l != 0 {
					bb += ")"
				} else {
					if f == 0 {
						leftNode := backTracking(bb)
						h.Left = leftNode
						bb = ""
						f = 1
					} else {
						rightNode := backTracking(bb)
						h.Right = rightNode
					}
				}
			}
		}

		return h
	}

	return backTracking(s)
}
