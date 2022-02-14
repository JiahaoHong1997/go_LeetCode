package leetcode93

import "strconv"

func restoreIpAddresses(s string) []string {

	if len(s) > 12 || len(s) < 4 {
		return []string{}
	}

	ret := []string{}
	var backTracking func(string, int, string)
	backTracking = func(s1 string, count int, path string) {

		if count == 4 {
			if len(s1) <= 3 {
				x, _ := strconv.Atoi(s1)
				if (x >= 0 && x <= 9 && len(s1) == 1) || (x >= 10 && x <= 99 && len(s1) == 2) || (x >= 100 && x <= 255 && len(s1) == 3) {
					path += s1
					ret = append(ret, path)
				}
			}
			return
		}

		for i:=1; i<=3; i++ {
			if len(s1) < i {
				break
			}
			if s1[0] == '0' {
				backTracking(s1[1:], count+1, path+"0.")
				break
			}
			if i == 3 {
				x, _ := strconv.Atoi(s1[:3])
				if x <= 255 {
					backTracking(s1[3:], count+1, path+s1[:3]+".")
				}
				break
			}
			backTracking(s1[i:], count+1, path+s1[:i]+".")
		}
	}
	backTracking(s, 1, "")
	return ret
}