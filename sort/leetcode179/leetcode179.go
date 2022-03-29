package leetcode179

import (
	"strconv"
	"sort"
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

	ret := ""
	for i:=0; i<len(nums); i++ {
		ret += strconv.Itoa(nums[i])
	}
	return ret
}