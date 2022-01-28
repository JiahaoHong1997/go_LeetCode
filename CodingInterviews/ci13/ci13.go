package ci13

func movingCount(m int, n int, k int) int {
	traveled := make(map[int]bool)
	count := 0

	var countBitSum func(i,j int) bool
	countBitSum = func(i, j int) bool {
		sum := 0
		for i>0 || j>0 {
			var x int
			if i != 0 {
				x = i % 10
				sum += x
				i /= 10
			} 
			if j != 0 {
				x = j % 10
				sum += x
				j /= 10	
			}
		}
		if sum > k {
			return false
		}
		return true
	}

	h_idx := []int{0,1,-1,0}
	w_idx := []int{1,0,0,-1}

	var goodRegion func(int, int) 
	goodRegion = func(h, w int) {
		traveled[h*n+w] = true
		if !countBitSum(h,w) {
			return
		} else {
            // fmt.Println(h,w)
			count++
		}

		for k:=0; k<4; k++ {
			mh, mw := h+h_idx[k], w+w_idx[k]
			if mh >= 0 && mh < m && mw >= 0 && mw < n && !traveled[mh*n+mw] {
				goodRegion(mh, mw)
			}
		}

	}
    goodRegion(0,0)
	return count
}