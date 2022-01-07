package leetcode254


func getFactors(n int) [][]int {

	ret := make([][]int,0)
	l := make([]int,0)

	var backTracking func(int)
	backTracking = func(k int) {
		if k <= 1 {
			return
		}

		for i:=2; i*i<=k; i++ {
			if len(l) > 0 {  // 避免重复选择
				if i < l[len(l)-1] || k/i < l[len(l)-1] {
					continue
				}
			}

			if k%i == 0 {
				l = append(l, i, k/i)
				l1 := make([]int,len(l))
				copy(l1,l)
				ret = append(ret, l1)

				l = l[:len(l)-1]
				backTracking(k/i)
				l = l[:len(l)-1]
			}
		}
	}

	backTracking(n)
	return ret
}