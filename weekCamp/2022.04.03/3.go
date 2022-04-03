package weekCamp

func maximumCandies(candies []int, k int64) int {
	var total int64
	for i:=0; i<len(candies); i++ {
		total += int64(candies[i])
	}

	if total < k {
		return 0
	} else if total == k {
		return 1
	}

	var satisified func(int) bool
	satisified = func(n int) bool {
		var count int64
		for i:=0; i<len(candies); i++ {
			count += int64(candies[i]/n)
			if count >= k {
				return true
			}
		}
		return false
	}
    
    
    
	l, r := 1, int(total/k)
    if satisified(r) {
        return r
    }
	for l < r {
		m := l+(r-l)>>1
		if satisified(m) {
			l = m+1
		} else if !satisified(m) {
			r = m
		}
	}

	return l-1
}