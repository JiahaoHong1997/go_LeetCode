/*
 * @Descripttion: 有效的括号
 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。

示例 1：
输入：s = "()"
输出：true

示例 2：
输入：s = "()[]{}"
输出：true

示例 3：
输入：s = "(]"
输出：false
 
提示：
1 <= s.length <= 10^4
s 仅由括号 '()[]{}' 组成
 * @Solution: 用栈来做
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-09-12 21:47:49
 * @LastEditors: 洪笳淏
 */
package stack

func isValid(s string) bool {
    stack := []byte{}

    for i:=0; i<len(s); i++ {
        switch s[i] {
        case '(','{','[':
            stack = append(stack,s[i])
        default:
            if len(stack) == 0 {
                return false
            }
            if (s[i] == ')' && stack[len(stack)-1] == '(') || (s[i] == '}' && stack[len(stack)-1] == '{') || (s[i] == ']' && stack[len(stack)-1] == '[') {
                stack = stack[:len(stack)-1]
            } else {
                return false
            }
        }
    }

    if len(stack) == 0 {
        return true
    }

    return false
}