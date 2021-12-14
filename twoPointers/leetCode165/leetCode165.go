package leetcode165

import (
	"strings"
	"strconv"
)

func compareVersion(version1 string, version2 string) int {
	str1 := strings.Split(version1, ".")
	str2 := strings.Split(version2, ".")

	n1, n2 := len(str1), len(str2)

	for i, j := 0, 0; i < n1 || j < n2; {

		if i == n1 {
			for _, v := range str2[j] {
				if v != '0' {
					return -1
				}
			}
			j++
			continue
		} else if j == n2 {
			for _, v := range str1[i] {
				if v != '0' {
					return 1
				}
			}
			i++
			continue
		}

		s1, s2 := str1[i], str2[j]
		if s1 == s2 {
			i++
			j++
			continue
		}

		
		b1, b2 := []byte(s1), []byte(s2)
		for i:=0; i<len(b1); {
			if b1[i] == '0' {
				b1 = b1[1:]
			} else {
				break
			}
		}
		for i:=0; i<len(b2); {
			if b2[i] == '0' {
				b2 = b2[1:]
			} else {
				break
			}
		}
		if len(b1) == 0 && len(b2) != 0 {
			return -1
		} else if len(b1) != 0 && len(b2) == 0 {
			return 1
		} else if len(b1) == 0 && len(b2) == 0 {
			i++
			j++
			continue
		}
		t1, _ := strconv.Atoi(string(b1))
		t2, _ := strconv.Atoi(string(b2))
		if t1 == t2 {
			i++
			j++
			continue
		} else if t1 > t2 {
			return 1
		} else {
			return -1
		}
	}
	return 0
}
