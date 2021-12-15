package leetcode424

func characterReplacement(s string, k int) int {
	window := make(map[byte]int)
	l, r := 0, 0
	maxNumInWindow := 0
	length, ret := 0, 0

	for r < len(s) {
		x := s[r]
		r++

		length++
		window[x]++
		maxNumInWindow = max(maxNumInWindow,window[x])
		if length-maxNumInWindow <= k {
			ret = max(ret,length)
			continue
		}

		for length-maxNumInWindow > k {
			remove := s[l]
			length--
			l++
			if window[remove] != maxNumInWindow {
				if window[remove] == 1 {
					delete(window,remove)
				} else {
					window[remove]--
				}
				if length-maxNumInWindow <= k {
					break
				}
			} else {
				window[remove]--
				tmp := 0
				for _,v := range window {
					tmp = max(tmp,v)
				}
				maxNumInWindow = tmp
			}
		}

	}
	return ret
}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}