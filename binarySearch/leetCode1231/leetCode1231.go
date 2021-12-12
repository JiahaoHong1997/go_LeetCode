package leetcode1231

import "math"

func maximizeSweetness(sweetness []int, k int) int {
	var (
		left = math.MaxInt32
		right = 0
	)

	for _, v := range sweetness {
		left = min(left, v)
		right += v
	}

	ret := 0
	for left < right {
		mid := left + (right-left)/2
		val := check(sweetness, mid, k+1)
		if val[0] <= 0 {
			left = mid + 1
			ret = val[1]
		} else {
			right = mid
		}
	}
	return ret
}

func check(nums []int, mid int, k int) []int {

	minpicess := math.MaxInt32
	sum := 0
	for i:=0; i<len(nums); i++ {
		sum += nums[i]

		if sum >= mid {
			k--
			minpicess = min(minpicess, sum)
			sum = 0
		}
	}
	return []int{k,minpicess}
}

func min(a,b int) int {
	if a<b {
		return a
	}
	return b
}