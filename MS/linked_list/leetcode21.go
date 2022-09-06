/*
 * @Descripttion:
 * @Solution:
 * @version:
 * @Author: 洪笳淏
 * @Date: 2022-09-06 16:01:46
 * @LastEditors: 洪笳淏
 */
package linkedlist

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    head := &ListNode{Val:0}
    p := &ListNode{Val:-101}

    for list1 != nil || list2 != nil {
        if list1 == nil {
            if p.Val == -101 {
                return list2
            } else {
                p.Next = list2
                return head.Next
            }
        } else if list2 == nil {
            if p.Val == -101 {
                return list1
            } else {
                p.Next = list1
                return head.Next
            }
        }

        if list1.Val < list2.Val {
            if head.Next == nil {
                p = list1
                head.Next = p
            } else {
                p.Next = list1
                p = p.Next
            }
            list1 = list1.Next
            p.Next = nil
        } else {
            if head.Next == nil {
                p = list2
                head.Next = p
            } else {
                p.Next = list2
                p = p.Next
            }
            list2 = list2.Next
            p.Next = nil
        }
    }

    return head.Next
}