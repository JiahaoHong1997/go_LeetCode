package leetcode16

import (
	"sort"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) == 3 {
		return nums[0]+nums[1]+nums[2]
	}
	sort.Ints(nums)
	ret := math.MaxInt32

	for i:=0; i<len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			x := sum(nums[i],nums[l],nums[r]) 
			if x == target {
				return target
			} else {
				if abs(ret,target) > abs(x,target) {
                    ret = x
                }
				if x > target {
					r--
				} else {
					l++
				}
			}
			
		}
	
	}
	return ret
}

func abs(a,b int) int {
	if a >= b {
		return a-b
	}
	return b-a
}

func sum(a,b,c int) int {
	return a+b+c
}