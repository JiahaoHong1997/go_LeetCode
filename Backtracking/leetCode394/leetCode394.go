package leetcode394

import "strconv"


func decodeString(s string) string {
    ret := ""
	
	var helper func(*int) (string,int)
	helper = func(i *int) (string,int) {

		s1 := ""
		kString := ""
		k := 0
        count := 0
		for ; *i < len(s); *i++ {
			if int(s[*i]) - '0' >= 0 && int(s[*i] - '0') <= 9 && count == 0 { // 最外层的数字累计
				kString += string(s[*i])
			} else if int(s[*i]) - '0' >= 0 && int(s[*i] - '0') <= 9 && count > 0 {		// 出现内层，递归调用来处理
                
				s2, k2 := helper(i)
                // fmt.Println(k2,*i)
				for ; k2 >0; k2-- {
					s1 += s2
				}
			} else if s[*i] == '[' && count == 0 {	// 最外层重复次数计算
                k, _ = strconv.Atoi(kString)
                // fmt.Println(k,*i)
				count++
				continue
			} else if s[*i] == ']' {	// 最外层结束，出栈
				break
			} else {					// 最外层字符串
				s1 += string(s[*i])
                // fmt.Println(s1)
			}

			
		}
		return s1, k
	}

	for i:=0; i<len(s); i++ {
		if int(s[i]) - '0' >= 1 && int(s[i] - '0') <= 9 {
			s1, k := helper(&i)
			for ; k>0; k-- {
				ret += s1
			}
            // fmt.Println(i,ret,s1)
		} else if s[i] != ']' {
			ret += string(s[i])
		}

	}
	return ret
}