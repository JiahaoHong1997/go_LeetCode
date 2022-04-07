package leetcode163

import "strconv"

func findMissingRanges(nums []int, lower int, upper int) []string {

	ret := []string{}
	if len(nums) == 0 {
		if lower == upper {
			s := strconv.Itoa(lower)
			ret = append(ret, s)
		} else {
			s := strconv.Itoa(lower) + "->" + strconv.Itoa(upper)
			ret = append(ret, s)
	}
		return ret
	}

	pre := lower
	for i:=0; i<len(nums); i++ {
		var s string
		if i == 0 {
			if nums[i] != pre {
				if nums[i] == pre+1 {
					s = strconv.Itoa(pre)
				} else {
					s = strconv.Itoa(pre) + "->" + strconv.Itoa(nums[i]-1)
				}
				ret = append(ret, s)
				pre = nums[i]
			}
		} else {
			if nums[i] == pre + 2 {
				s = strconv.Itoa(nums[i]-1)
			} else if nums[i] == pre + 1 {
				s = ""
			} else {
				s = strconv.Itoa(pre+1) + "->" + strconv.Itoa(nums[i]-1)
			}

			pre = nums[i]
			if s != "" {
				ret = append(ret, s)
			}
		}
		
		if i == len(nums)-1 {
			if nums[i] != upper {
				if nums[i] == upper-1 {
					s = strconv.Itoa(upper)
				} else {
					s = strconv.Itoa(nums[i]+1) + "->" + strconv.Itoa(upper)
				}
				ret = append(ret, s)
			}
			break
		}
	}

	return ret
}