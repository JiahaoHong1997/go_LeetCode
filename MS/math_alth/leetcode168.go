/*
 * @Descripttion: 从 1 开始的 26 进制
 给你一个整数 columnNumber ，返回它在 Excel 表中相对应的列名称。

例如：
A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28 
...

示例 1：
输入：columnNumber = 1
输出："A"

示例 2：
输入：columnNumber = 28
输出："AB"

示例 3：
输入：columnNumber = 701
输出："ZY"

示例 4：
输入：columnNumber = 2147483647
输出："FXSHRXW"
 * @Solution:
这是一道从 1 开始的的 26 进制转换题。
对于一般性的进制转换题目，只需要不断地对 columnNumber 进行 % 运算取得最后一位，然后对 columnNumber 进行 / 运算，将已经取得的位数去掉，直到 columnNumber 为 0 即可。
一般性的进制转换题目无须进行额外操作，是因为我们是在「每一位数值范围在 [0,x)」的前提下进行「逢 x 进一」。
但本题需要我们将从 1 开始，因此在执行「进制转换」操作前，我们需要先对 columnNumber 执行减一操作，从而实现整体偏移。

 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-08-27 16:14:43
 * @LastEditors: 洪笳淏
 */
package mathalth

func convertToTitle(columnNumber int) string {
    
    res := ""
    for columnNumber > 0 {
        columnNumber--
        res = string('A'+columnNumber%26) + res
        columnNumber /= 26
    }
    return res

}