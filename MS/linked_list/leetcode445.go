/*
 * @Descripttion: 两数相加 II
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
你可以假设除了数字 0 之外，这两个数字都不会以零开头。

示例1：
输入：l1 = [7,2,4,3], l2 = [5,6,4]
输出：[7,8,0,7]

示例2：
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[8,0,7]

示例3：
输入：l1 = [0], l2 = [0]
输出：[0]
 
提示：
链表的长度范围为 [1, 100]
0 <= node.val <= 9
输入数据保证链表代表的数字无前导 0
 * @Solution: 方法一：使用栈  方法二：反转两个链表后再做
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-09-15 13:22:53
 * @LastEditors: 洪笳淏
 */
package linkedlist

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    stack1 := make([]int,0)
    stack2 := make([]int,0)

    for l1 != nil || l2 != nil {
        if l1 != nil {
            stack1 = append(stack1, l1.Val)
            l1 = l1.Next
        }

        if l2 != nil {
            stack2 = append(stack2, l2.Val)
            l2 = l2.Next
        }
    }

    var p *ListNode
    c := 0
    for len(stack1) != 0 || len(stack2) != 0 {
        n1, n2 := len(stack1), len(stack2)
        dummy := &ListNode{Next:nil}
        if n1 == 0 {
            dummy = &ListNode{Val:(stack2[n2-1]+c)%10, Next:nil}
            c = (stack2[n2-1]+c)/10
            stack2 = stack2[:n2-1]
        } else if n2 == 0 {
            dummy = &ListNode{Val:(stack1[n1-1]+c)%10, Next:nil}
            c = (stack1[n1-1]+c)/10
            stack1 = stack1[:n1-1]
        } else {
            val := (stack1[n1-1]+stack2[n2-1]+c)%10
            c = (stack1[n1-1]+stack2[n2-1]+c)/10
            dummy = &ListNode{Val:val, Next:nil}
            stack1 = stack1[:n1-1]
            stack2 = stack2[:n2-1]
        }
        dummy.Next = p 
        p = dummy
    }

    if c == 1 {
        dummy := &ListNode{Val:1, Next:p}
        p = dummy
    }
    return p
}