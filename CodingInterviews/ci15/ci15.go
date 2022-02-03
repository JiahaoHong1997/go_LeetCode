package ci15

func hammingWeight(num uint32) int {
	count := 0

	for num > 0 {
		if num ^ (num-1) == 1 {
			count++
		}
		num >>= 1
	}

	return count
}