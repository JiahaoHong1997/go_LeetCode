package leetcode415

func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	b := 0
	s := ""
	for i >= 0 || j >= 0 {
		if i < 0 {
			if b == 0 {
				s = num2[:j+1] + s
				return s
			}
			x := int(num2[j]) - '0' + 1
			if x > 9 {
				b = 1
				a := x % 10
				s = string(a + '0') + s
			} else {
                b = 0
			    s = string(x + '0') + s
            }
			j--
			continue
		}

		if j < 0 {
			if b == 0 {
				s = num1[:i+1] + s
				return s
			}
			x := int(num1[i]) - '0' + 1
			if x > 9 {
				b = 1
				a := x % 10
				s = string(a + '0') + s
			} else {
			    b = 0
			    s = string(x + '0') + s
            }
			i--
			continue
		}

		x1, x2 := num1[i], num2[j]
		x := (int(x1) - '0') + (int(x2) - '0') + b
		a := x % 10 + '0'
		b = x / 10
        
		s = string(a) + s
		i--
		j--
	}

    if b == 1 {
		s = "1" + s
	}
	return s
}
