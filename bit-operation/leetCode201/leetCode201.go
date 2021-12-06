package leetcode201

func rangeBitwiseAnd(left int, right int) int {
	ret := left
	for i:=left; i<=right; i++ {
		ret = ret & i
		if ret == 0 {
			return 0
		}
	}
	return ret
}