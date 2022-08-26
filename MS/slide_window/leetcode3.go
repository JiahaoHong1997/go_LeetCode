/*
 * @Descripttion: 无重复字符的最长子串
 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:
输入: s = "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。 

提示：
0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成

 * @Solution: 滑动窗口基础题
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-08-26 12:25:22
 * @LastEditors: 洪笳淏
 */
package slidewindow

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}

	record := make(map[byte]int, 0)
	l, r := 0, 0
	res := 0
	for r < len(s) {
		x := s[r]
		r++
		record[x]++

		for len(record) < r-l {
			y := s[l]
			l++
			record[y]--
			if record[y] == 0 {
				delete(record, y)
			}
		}
		res = max(res, r-l)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
