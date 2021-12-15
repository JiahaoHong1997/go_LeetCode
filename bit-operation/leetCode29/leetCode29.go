package leetcode29

import "math"

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	lessThanZero := false
	if dividend < 0 && divisor < 0 {
		dividend = 0 - dividend
		divisor = 0 - divisor
	} else if dividend < 0 {
		dividend = 0 - dividend
		lessThanZero = true
	} else if divisor < 0 {
		divisor = 0 - divisor
		lessThanZero = true
	}

	if dividend < divisor {
		return 0
	}

	a,b := dividend, divisor
	ret := 0
	for a >= b {
		base := b 
		cnt := 1

		for a >= (base<<1) {
			base <<= 1
			cnt <<= 1
		}
		ret += cnt
		a -= base
	}

	if ret >= int(math.Pow(2, 31)) && !lessThanZero {
		ret = int(math.Pow(2, 31)) - 1
	} else if ret > int(math.Pow(2, 31)) && lessThanZero {
        ret = int(math.Pow(2, 31))
    }

	if lessThanZero {
		return 0 - ret
	}
	return ret
}
