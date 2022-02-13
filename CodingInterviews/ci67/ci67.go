package ci67

import "math"

func strToInt(str string) int {

	if str == "" {
		return 0
	}

	var atoi func(string) int 
	atoi = func(s string) int {

		if s == "" {
			return 0
		}
		
		x := 0
		if s[0] == '-' || s[0] == '+' {
			for j:=1; j<len(s); j++ {
				x *= 10
				x += int(s[j])-'0'
				if x > math.MaxInt32 && s[0] == '-' {
					return math.MinInt32
				} else if x >= math.MaxInt32 && s[0] == '+' {
					return math.MaxInt32
				}
			}
			if s[0] == '-' {
				return -1*x
			} else {
				return x
			}
		}

		for j:=0; j<len(s); j++ {
			x *= 10
			x += int(s[j])-'0'
			if x >= math.MaxInt32 {
				return math.MaxInt32
			}
		}
		return x
	}


	var checkLegal func() int
	checkLegal = func() int {

		isLegel := false
		s := ""
		for i:=0; i<len(str); i++ {
			if !isLegel && str[i] == ' ' {
				continue
			} else if !isLegel {	
				if str[i] >= '0' && str[i] <= '9' {
					isLegel = true
					s += str[i:i+1]
				} else if (str[i] == '-' || str[i] == '+') && i < len(str)-1 && str[i+1] >= '0' && str[i+1] <= '9' {
					isLegel = true
					s += str[i:i+1]
				} else {
					return 0
				}
			} else if str[i] >= '0' && str[i] <= '9' {
				s += str[i:i+1]
			} else {
				break
			}
		}		
		return atoi(s)
	}

	x := checkLegal()
	if x == 0 {
		return 0
	}

	return x
}