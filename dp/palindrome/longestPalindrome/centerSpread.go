package longestpalindrome

func longestpalindromeCS(s string) string {
	if s == "" {
		return ""
	}

	begin, end := 0, 0
	for i:=0; i<len(s); i++ {
		left1, right1 := expandAroundCenter(s,i,i)
		left2, right2 := expandAroundCenter(s,i,i+1)

		if right1-left1 > end-begin {
			begin, end = left1, right1
		}

		if right2-left2 > end-begin {
			begin, end = left2, right2
		}
	}
	return s[begin:end+1]
}

func expandAroundCenter(s string, left int, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; {
		left--
		right++
	}

	return left+1, right-1
}