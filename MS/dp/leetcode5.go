/*
 * @Descripttion: 最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。
示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

示例 2：
输入：s = "cbbd"
输出："bb"

提示：
1 <= s.length <= 1000
s 仅由数字和英文字母组成
 * @Solution: 中心扩散 / 传统动态规划
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-09-06 17:43:50
 * @LastEditors: 洪笳淏
 */
package dp

func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		left1, right1 := centerExpand(s, i, i)
		left2, right2 := centerExpand(s, i, i+1)
		if right1-left1 >= len(res) {
			res = s[left1 : right1+1]
		}
		if right2-left2 >= len(res) {
			res = s[left2 : right2+1]
		}
	}

	return res
}

func centerExpand(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return left + 1, right - 1
}
