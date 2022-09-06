/*
 * @Descripttion: 合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2：
输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
 

提示：
1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104
 * @Solution: 先对二维数组排序，之后再逐个判断是否要和res融合
 * @version: 1.0
 * @Author: 洪笳淏
 * @Date: 2022-09-06 17:54:54
 * @LastEditors: 洪笳淏
 */
package sort

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	for i := 0; i < len(intervals); i++ {
		left, right := intervals[i][0], intervals[i][1]
		if len(res) == 0 {
			res = append(res, intervals[i])
			continue
		}
		lastArr := res[len(res)-1]
		lastRight := lastArr[1]
		if left > lastRight {
			res = append(res, intervals[i])
		} else if right <= lastRight {
			continue
		} else {
			lastArr[1] = right
		}
	}

	return res
}
