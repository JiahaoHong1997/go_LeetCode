/*
 * @Descripttion: 字符串相加
 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。
你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。

示例 1：
输入：num1 = "11", num2 = "123"
输出："134"

示例 2：
输入：num1 = "456", num2 = "77"
输出："533"

示例 3：
输入：num1 = "0", num2 = "0"
输出："0"
 
提示：
1 <= num1.length, num2.length <= 10^4
num1 和num2 都只包含数字 0-9
num1 和num2 都不包含任何前导零
 * @Solution:
 * @version:
 * @Author: 洪笳淏
 * @Date: 2022-09-12 22:54:53
 * @LastEditors: 洪笳淏
 */
package simulation

import (
	"strconv"
)

func addStrings(num1 string, num2 string) string {
    var res string
    var c int
    for len(num1) > 0 || len(num2) > 0 {
        if len(num2) == 0 {
            if c == 0 {
                res = num1 + res
                return res
            } else {
                newBit := (int(num1[len(num1)-1])-'0'+c)%10
                c = (int(num1[len(num1)-1])-'0'+c)/10
                res = strconv.Itoa(newBit)+res
                num1 = num1[:len(num1)-1]
            }
        } else if len(num1) == 0 {
            if c == 0 {
                res = num2 + res
                return res
            } else {
                newBit := (int(num2[len(num2)-1])-'0'+c)%10
                c = (int(num2[len(num2)-1])-'0'+c)/10
                res = strconv.Itoa(newBit)+res
                num2 = num2[:len(num2)-1]
            }
        } else {
            newBit := ((int(num1[len(num1)-1])-'0')+(int(num2[len(num2)-1])-'0')+c)%10
            c = ((int(num1[len(num1)-1])-'0')+(int(num2[len(num2)-1])-'0')+c)/10
            res = strconv.Itoa(newBit)+res
            num1 = num1[:len(num1)-1]
            num2 = num2[:len(num2)-1]
        }
    }
    if c != 0 {
        res = "1" + res
    }
    return res
}