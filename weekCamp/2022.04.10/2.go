package weekCamp

import (
	"strconv"
	"strings"
)

func minimizeResult(expression string) string {
    
	var cal func(string) int 
	cal = func(s string) int {
		s1 := ""
		i := 0
		for ; i<len(s); i++ {
			if s[i] >= '0' && s[i] <= '9' {
				s1 += s[i:i+1]
			} else {
				break
			}
		}
		
		i++
		a, b, pre := "", "", true
		for ; i<len(s); i++ {
			if s[i] >= '0' && s[i] <= '9' && pre {
				a += s[i:i+1]
			} else if s[i] == '+' {
				pre = false
			} else if s[i] >= '0' && s[i] <= '9' && !pre {
				b += s[i:i+1]
			} else if s[i] == ')' {
				break
			}
		}
		
        var preNum int
        if s1 == "" {
            preNum = 1  
        } else {
            preNum, _ = strconv.Atoi(s1) 
        }
		
		mid1, _ := strconv.Atoi(a)
		mid2, _ := strconv.Atoi(b)
		
		s2 := ""
		for ; i<len(s); i++ {
			if s[i] >= '0' && s[i] <= '9' {
				s2 += s[i:i+1]
			} 
		}
		if s2 != "" {
			lastNum, _ := strconv.Atoi(s2)
			return preNum*(mid1+mid2)*lastNum
		} 
		
		return preNum*(mid1+mid2)
	}

	str := strings.Split(expression, "+")
	x1, _ := strconv.Atoi((str[0]))
	x2, _ := strconv.Atoi(str[1])
	ret := x1+x2
    retS := "("+expression+")"

	b1 := []byte(str[0])
	b2 := []byte(str[1])

	for i:=0; i<len(b1); i++ {
		b1 = append(b1[:i], append([]byte{'('}, b1[i:]...)...)
        // fmt.Println(b1)
		for j:=1; j<=len(b2); j++ {
			b2 = append(b2[:j], append([]byte{')'}, b2[j:]...)...)
			t := cal(string(b1)+"+"+string(b2))
            //fmt.Println(string(b1)+"+"+string(b2), t)
			if t < ret {
				ret = t
				retS = string(b1)+"+"+string(b2)
			}
			b2 = append(b2[:j], b2[j+1:]...)
		}
		b1 = append(b1[:i], b1[i+1:]...)
	}

	return retS
}
