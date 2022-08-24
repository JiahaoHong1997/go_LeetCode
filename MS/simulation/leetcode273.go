/**
 * @Descripttion:
 将非负整数 num 转换为其对应的英文表示。
示例 1：
输入：num = 123
输出："One Hundred Twenty Three"

示例 2：
输入：num = 12345
输出："Twelve Thousand Three Hundred Forty Five"

示例 3：
输入：num = 1234567
输出："One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"

提示：
0 <= num <= 2^31 - 1
 * @Solution: 模拟题
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-08-24 16:00:12
 * @LastEditors: 洪笳淏
*/

package simulation

const (
	hundred  = "Hundred"
	thousand = "Thousand"
	million  = "Million"
	billion  = "Billion"
	ty       = "ty"
)

var numWord map[int]string

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	numWord = map[int]string{
		1:  "One",
		2:  "Two",
		3:  "Three",
		4:  "Four",
		5:  "Five",
		6:  "Six",
		7:  "Seven",
		8:  "Eight",
		9:  "Nine",
		10: "Ten",
		11: "Eleven",
		12: "Twelve",
		13: "Thirteen",
		14: "Fourteen",
		15: "Fifteen",
		16: "Sixteen",
		17: "Seventeen",
		18: "Eighteen",
		19: "Nineteen",
		20: "Twenty",
		30: "Thirty",
		40: "Forty",
		50: "Fifty",
		60: "Sixty",
		70: "Seventy",
		80: "Eighty",
		90: "Ninety",
	}
	res := ""
	if num/1000000000 > 0 {
		tmp := num / 1000000000
		num = num % 1000000000
		res += numWord[tmp] + " " + billion + " "
	}
	if num/1000000 > 0 {
		tmp := num / 1000000
		num = num % 1000000
		res += lessThanThousand(tmp) + million + " "
	}
	if num/1000 > 0 {
		tmp := num / 1000
		num = num % 1000
		res += lessThanThousand(tmp) + thousand + " "
	}
	if num > 0 {
		res += lessThanThousand(num)
	}
	length := len(res)
	return res[:length-1]
}

func lessThanThousand(num int) string {
	res := ""
	if num/1000 > 0 {
		tmp := num / 1000
		num = num % 1000
		res += numWord[tmp] + " " + thousand + " "
	}
	if num/100 > 0 {
		tmp := num / 100
		num = num % 100
		res += numWord[tmp] + " " + hundred + " "
	}
	if num/10 > 0 {
		if num/10 == 1 {
			res += numWord[num] + " "
			return res
		}
		tmp := num / 10 * 10
		num = num % 10
		res += numWord[tmp] + " "
	}
	if num > 0 {
		res += numWord[num] + " "
	}
	return res
}
