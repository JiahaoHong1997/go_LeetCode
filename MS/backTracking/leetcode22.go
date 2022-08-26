/*
 * @Descripttion: 括号生成
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

示例 2：
输入：n = 1
输出：["()"]

提示：
1 <= n <= 8
 * @Solution:
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-08-26 14:20:21
 * @LastEditors: 洪笳淏
*/
package backtracking

func generateParenthesis(n int) []string {
	var res []string
	var b []byte
	var backTracking func(int, int, []byte)
	backTracking = func(left, right int, str []byte) {
		if left == n && right == n {
			tmp := make([]byte, len(str))
			copy(tmp, str)
			res = append(res, string(tmp))
			return
		}

		if left < n {
			str = append(str, '(')
			backTracking(left+1, right, str)
			str = str[:len(str)-1]
		}
		if left > right {
			str = append(str, ')')
			backTracking(left, right+1, str)
			str = str[:len(str)-1]
		}
	}

	backTracking(0, 0, b)
	return res
}
