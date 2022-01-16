package leetcode301

import "math"

func removeInvalidParentheses(s string) []string {

	path := []string{}
	b := []byte(s)
    m := make(map[string]bool)
	count := math.MaxInt32

	var backTracking func(int, int, int, int)
	backTracking = func(p int, r int, left int, right int) {

		if r > count {
			return
		}

		if p == len(b) {
			if left == right {
				t := make([]byte, len(b))
				copy(t, b)

				if count > r {
					path = []string{}
					count = r
				}
                if !m[string(t)] {
                    path = append(path, string(t))
                    m[string(t)] = true
                }
				
			}
			return
		}

		x := b[p]
		if x == '(' {
			backTracking(p+1, r, left+1, right)

			b = append(b[:p], b[p+1:]...)
			backTracking(p, r+1, left, right)
			b = append(b[:p], append([]byte{x}, b[p:]...)...)
		} else if x == ')' {

			if left > right {
				backTracking(p+1, r, left, right+1)
			}

			b = append(b[:p], b[p+1:]...)
			backTracking(p, r+1, left, right)
			b = append(b[:p], append([]byte{x}, b[p:]...)...)

		} else {
			backTracking(p+1, r, left, right)
		}

	} 

    backTracking(0, 0, 0, 0)
	return path
}