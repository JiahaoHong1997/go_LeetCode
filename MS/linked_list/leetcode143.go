/*
 * @Descripttion: 重排链表
 给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1：
输入：head = [1,2,3,4]
输出：[1,4,2,3]

示例 2：
输入：head = [1,2,3,4,5]
输出：[1,5,2,4,3]
 

提示：
链表的长度范围为 [1, 5 * 10^4]
1 <= node.val <= 1000
 * @Solution: 先将链表的后半部分反转，然后再间隔插入
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-09-16 17:30:23
 * @LastEditors: 洪笳淏
 */
package linkedlist

func reorderList(head *ListNode) {
	fast, slow := head, head
	var tmp *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next

		if fast == nil {
			tmp = slow.Next
			slow.Next = nil
			break
		}
		slow = slow.Next
	}

	var secondHead *ListNode
	if fast == nil {
		secondHead = reverseLinkedList(tmp)
	} else {
		tmp = slow.Next
		slow.Next = nil
		secondHead = reverseLinkedList(tmp)
	}

	opHead := &ListNode{Next: head}
	var count int
	for opHead != nil && secondHead != nil {
		if count%2 == 0 {
			opHead = opHead.Next
		} else {
			next := opHead.Next
			opHead.Next = secondHead
			tmp := secondHead.Next
			secondHead.Next = next
			secondHead = tmp
			opHead = opHead.Next
		}
		count++
	}
}

func reverseLinkedList(p *ListNode) *ListNode {
	if p == nil || p.Next == nil {
		return p
	}

	pre, cur := p, p.Next
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	p.Next = nil
	return pre
}
