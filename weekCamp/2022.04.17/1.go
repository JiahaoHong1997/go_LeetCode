package weekCamp

import "strconv"

func digitSum(s string, k int) string {

	if len(s) <= k {
		return s
	}

	for len(s) > k {

		t := ""
		for i:=0; i<=(len(s)/k)*k; i=i+k {
			x := ""
			if i == (len(s)/k)*k {
				x = s[i:]
			} else {
				x = s[i:i+k]
			}

			if x == "" {
				break
			}

			var n int
			for j:=0; j<len(x); j++ {
				n += int(x[j])-'0'
			}

			t += strconv.Itoa(n)
		}
		s = t
	}

	return s
}
