package leetcode287

func findDuplicateBS(nums []int) int {
	n := len(nums)-1
	l, r := 1, n

	var repeat int
	for l < r {
		mid := l+(r-l)/2
		repeat = searchLess(nums, mid)
		if repeat <= mid {
			l = mid + 1
		} else {
			r = mid
		}
	}


	return l
}

func searchLess(nums []int, target int) int {
	ret := 0
	for _, v := range nums {
		if v <= target {
			ret++
		}
	}
	return ret
}