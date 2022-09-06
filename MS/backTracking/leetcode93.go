/*
 * @Descripttion: 复原 IP 地址
有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。

示例 1：
输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]

示例 2：
输入：s = "0000"
输出：["0.0.0.0"]

示例 3：
输入：s = "101023"
输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

提示：
1 <= s.length <= 20
s 仅由数字组成
 * @Solution:
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-09-06 17:32:11
 * @LastEditors: 洪笳淏
*/
package backtracking

import (
	"strconv"
)

func restoreIpAddresses(s string) []string {
	res := []string{}

	var backTracking func(string, string, int)
	backTracking = func(restS, path string, layer int) {
		if len(restS) == 0 {
			return
		}
		if layer == 4 {
			if restS[0] == '0' {
				if len(restS) == 1 {
					path += "0"
					tmp := make([]byte, len(path))
					copy(tmp, []byte(path))
					res = append(res, string(tmp))
				}
				return
			}
			restNum, _ := strconv.Atoi(restS)
			if restNum >= 0 && restNum <= 255 {
				path += restS
				tmp := make([]byte, len(path))
				copy(tmp, []byte(path))
				res = append(res, string(tmp))
			}
			return
		}

		for i := 1; i <= 3; i++ {
			if restS[0] == '0' {
				path += "0."
				backTracking(restS[1:], path, layer+1)
				path = path[:len(path)-2]
				break
			}
			if len(restS) < i {
				break
			}
			restNum, _ := strconv.Atoi(restS[:i])
			if restNum > 255 {
				break
			}
			path += restS[:i] + "."
			backTracking(restS[i:], path, layer+1)
			path = path[:len(path)-i-1]
		}
	}

	backTracking(s, "", 1)
	return res
}
