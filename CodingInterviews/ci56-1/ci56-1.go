package ci56

import "sort"

func singleNumbers(nums []int) []int {
	sort.Slice(nums, func(i,j int) bool {
		if nums[i] < nums[j] {
			return true
		}
		return false
	})
    
	x := nums[0]
	count := 1
	ret := []int{}
	for i:=1; i<len(nums); i++ {
        if i == len(nums)-1 {
            ret = append(ret, nums[i])
            return ret
        }
		if count == 1 {
			if nums[i] == x {
				count = 0
			} else {
				ret = append(ret, x)
				if len(ret) == 2 {
					return ret
				}
				x = nums[i]
			}
		} else {
			count = 1
			x = nums[i]
		}
	}

	return []int{}
}