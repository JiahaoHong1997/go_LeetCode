package leetcode340


func lengthOfLongestSubstringKDistinct(s string, k int) int {
	window := make(map[byte]int)
	l,r := 0, 0
	length, ret := 0, 0

	for r < len(s) {
		x := s[r]
		r++

		window[x]++
		length++
		if len(window) <= k {
			ret = max(ret,length)
			continue
		}

		for len(window) > k {
			remove := s[l]
			if window[remove] == 1 {
				delete(window,remove)
			} else {
				window[remove]--
			}
			length--
			l++
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