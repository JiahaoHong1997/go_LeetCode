package ci56

import "math"

func singleNumbersBit(nums []int) []int {
	x := nums[0]
	for i:=1; i<len(nums); i++ {
		x ^= nums[i]
	}

	c := 0
	for x^(x-1) != 1 {
		x >>= 1
		c++
	}

	a, b := math.MinInt32, math.MinInt32
	for i:=0; i<len(nums); i++ {
		t := nums[i]
		t >>= c
		if t^(t-1) == 1 {
			if a == math.MinInt32 {
				a = nums[i]
			} else {
				a ^= nums[i]
			}
		} else {
			if b == math.MaxInt32 {
				b = nums[i]
			} else {
				b ^= nums[i]
			}
		}
	}

	return []int{a,b}
}