package ci45

import (
	"fmt"
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x := fmt.Sprintf("%d%d", nums[i], nums[j])
		y := fmt.Sprintf("%d%d", nums[j], nums[i])
		return x < y
	})

	ret := ""
	for i := 0; i < len(nums); i++ {
		ret += strconv.Itoa(nums[i])
	}
	return ret
}
