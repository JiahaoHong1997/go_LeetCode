package ci48

func lengthOfLongestSubstring(s string) int {

	var (
		l = 0
		r = 0
	)

	var max func(int, int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	m := make(map[byte]int)
	ret := 0

	for r < len(s) {
		x := s[r]
		r++
		m[x]++

		for m[x] > 1 {
			y := s[l]
			l++
			m[y]--
		}

		ret = max(ret, r-l)
	}
	return ret
}