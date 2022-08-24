/**
 * @Descripttion:
 给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
不允许修改 链表。

 * @Solution: 翻译成人话：1.判断链表中是否存在环；2.找出环的起始节点
 是否存在环：使用快慢指针，如果快慢指针相遇，则存在环；否则如果快指针走到链表结尾，则一定不存在环；
 找到起始节点：当快慢指针相遇后，将其中一个指针放到head处，两指针都向后遍历，它们再次相遇之处便是起始节点。
 * @version:
 * @Author: 洪笳淏
 * @Date: 2022-08-24 16:12:50
 * @LastEditors: 洪笳淏
*/

package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast, slow := head, head
	isRing := false
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			isRing = true
			break
		}
	}

	if !isRing {
		return nil
	}
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
