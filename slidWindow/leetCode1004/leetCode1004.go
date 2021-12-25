package leetcode1004

func longestOnes(nums []int, k int) int {
	count := 0
	l, r := 0, 0
	ret := 0

	for r < len(nums) {
		x := nums[r]
		r++

		if x == 0 {
			count++
		}
		if count <= k {
			ret = max(ret,r-l)	
		}

		for count > k {
			x1 := nums[l]
			l++
			if x1 == 0 {
				count--
			} 
		}
	}
	return ret
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}