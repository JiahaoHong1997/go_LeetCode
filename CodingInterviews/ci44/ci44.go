package ci44

import "math"

func findNthDigit(n int) int {

	var countDigitNum func(int) (int, int)
	countDigitNum = func(x int) (int, int) {

		t := x
		d := 0
		for t > 0 {
			t /= 10
			d++
		}

		ret := 0
		t = d
		for t - 1 > 0 {
			ret += (t-1)*9*int(math.Pow10(t-2))
			t--
		}
		ret += d*(x - int(math.Pow10(d-1)) + 1)
		return d, ret
	}

    var sequence func(int) []int 
    sequence = func(x int) []int {

        ret := []int{}
        for x > 0 {
            ret = append([]int{x%10}, ret...)
            x /= 10
        }
        return ret
    }

	l, r := 0, n
	for l <= r {
		mid := l + (r-l)/2
		d, count := countDigitNum(mid)
		if count - d + 1 <= n && count >= n {
            r := sequence(mid)
			return r[n-count+d-1]
		} else if count < n {
			l = mid + 1
		} else if count - d + 1 > n {
			r = mid - 1
		}
	}

	return 0
}