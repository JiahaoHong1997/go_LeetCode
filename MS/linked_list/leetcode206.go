/**
 * @Descripttion: 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
 * @Solution: 反转链表，没啥好说的
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-08-25 11:44:04
 * @LastEditors: 洪笳淏
 */

package linkedlist

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    pre, current := head, head.Next
    for current != nil {
        next := current.Next
        current.Next = pre
        pre = current
        current = next
    }
    head.Next = nil
    return pre
}
