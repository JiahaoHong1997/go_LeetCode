package leetcode12

func intToRoman(num int) string {
	m := map[int]string{
		1:"I",
		5:"V",
		10:"X",
		50:"L",
		100:"C",
		500:"D",
		1000:"M",
	}

	s := ""
	i := 1
	for num > 0 {
		x := num % 10
		if x != 0 {
			if x >= 1 && x <= 3 {
				for j := 0; j < x; j++ {
					s = m[1*i] + s
				}
			} else if x == 4 {
				s = m[1*i] + m[5*i] + s
			} else if x == 5 {
				s = m[5*i] + s
			} else if x >= 6 && x <= 8 {
				for j := 0; j < x - 5; j++ {
					s = m[1*i] + s
				}
				s = m[5*i] + s
			} else {
				s = m[1*i] + m[10*i] + s
			}
		}
		num /= 10
		i *= 10
	}
	return s
}