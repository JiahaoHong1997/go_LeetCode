package leetcode201

func rangeBitwiseAndOP(left int, right int) int {
	i := 0
	for left != right {
		left >>= 1
		right >>= 1
		i++
	}
	for i > 0 {
		left <<= 1
		i--
	}
	return left
}