package ci38

import "sort" 

func permutation(s string) []string {
	ret := []string{}
	path := make([]byte, 0)
	b := []byte(s)
	n := len(s)
	sort.Slice(b, func(i, j int) bool {
		if b[i] < b[j] {
			return true
		}
		return false
	})

	var backTracking func([]byte)
	backTracking = func(b []byte) {
		if len(path) == n {
			tmp := make([]byte, n)
			copy(tmp, path)
			ret = append(ret, string(tmp))
		}

		for i:=0; i<len(b); i++ {
			if i>0 && b[i] == b[i-1] {
                continue
            }
			x := b[i]
			path = append(path, x)
			b = append(b[:i], b[i+1:]...)
			backTracking(b)
			b = append(b[:i], append([]byte{x}, b[i:]...)...)
			path = path[:len(path)-1]
		}
	}

	backTracking(b)
	return ret
}