package leetcode159

func lengthOfLongestSubstringTwoDistinct(s string) int {
	window := make(map[byte]int)

	l, r := 0, 0
	length := 0
	ret := 0

	for r < len(s) {
		x := s[r]
		r++

		window[x]++
		length++
		if len(window) <= 2 {
			ret = max(ret,length)
			continue
		}

		for len(window) > 2 {
			remove := s[l]
			l++
			if window[remove] == 1 {
				delete(window,remove)		
			} else {
				window[remove]--
			}
			length--
		}
		ret = max(ret,length)
		
	}
	return ret
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}