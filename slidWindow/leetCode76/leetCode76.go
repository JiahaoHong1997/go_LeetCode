package leetcode76

import "math"

func minWindow(s string, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for _,v := range t {
		need[byte(v)]++
	}
	n := len(need)
	count := math.MaxInt32
	ret := ""

	var (
		valid = 0
		l = 0
		r = 0
	)

	for r < len(s) {
		x := s[r]
		r++

		window[x]++
		if need[x] > 0 && window[x] == need[x] {
			valid++
		} 
		if valid == n {
			if r-l < count {
				ret = s[l:r]
				count = r-l
			}
		}

		for valid == n {
			x1 := s[l]
			l++
			if need[x1] > 0 && window[x1] == need[x1] {
				valid--
			}
			window[x1]--
			if window[x1] == 0 {
				delete(window, x1)
			}

			if valid == n-1 && r-l+1 < count {
				ret = s[l-1:r]
				count = r-l+1
			}
		}

	}
	return ret
}