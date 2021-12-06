package leetcode343

import "math"

func integerBreak(n int) int {
	ret := 0

	for i:=2; i<=n; i++ {
		if n % i == 0 {
			ret = max(ret, int(math.Pow(float64(n/i), float64(i))))
		} else {
			a := n-(n/i)*i
			b := i-a
			ret = max(ret, int(math.Pow(float64((n/i + 1)),float64(a)))*int(math.Pow(float64((n/i)),float64(b))))
		}
	}
	return ret
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}