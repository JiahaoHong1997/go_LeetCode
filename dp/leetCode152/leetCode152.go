package leetcode152

func maxProduct(nums []int) int {
	fmax, fmin := nums[0], nums[0]
	ret := nums[0]

    for i:=1; i<len(nums); i++ {
        x := nums[i]
        y := nums[i]*fmax
        z := nums[i]*fmin

        fmax = max(max(x,y),z)
        fmin = min(min(x,y),z)
        ret = max(ret, fmax)
    }


	return ret
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}