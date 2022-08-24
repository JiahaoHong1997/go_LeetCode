/**
 * @Descripttion: 
给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。

示例 1：
输入：nums = [10,2]
输出："210"

示例 2：
输入：nums = [3,30,34,5,9]
输出："9534330"
 * @Solution: 本题本质上是一个排序算法，不过排序规则不是单纯比较数字的实际大小，而是要逐位的比较。具体的做法就是A、B两数分别将对方置于自己之后，较大这应排序在前。
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-08-24 14:33:54
 * @LastEditors: 洪笳淏
 */
package sort

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
    sort.Slice(nums, func(i,j int)bool {
        x, y := nums[i], nums[j]
        sx, sy := 10, 10
        for sx <= x {
            sx *= 10
        }

        for sy <= y {
            sy *= 10
        }
        return sy*x+y > sx*y+x 
    })
    if nums[0] == 0 {
        return "0"
    }
    res := ""
    for _, num := range nums {
        res += strconv.Itoa(num)
    }
    return res
}