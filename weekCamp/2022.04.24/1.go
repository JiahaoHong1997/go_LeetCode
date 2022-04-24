package weekCamp

import "sort"

func intersection(nums [][]int) []int {

	m := make(map[int]int,0)
	n := len(nums)

	for i:=0; i<len(nums); i++ {
		for j:=0; j<len(nums[i]); j++ {
			m[nums[i][j]]++
		}
	}

	ret := []int{}
	for num, v := range m {
		if v == n {
			ret = append(ret, num)
		}
	}

	sort.Ints(ret)

	return ret
}