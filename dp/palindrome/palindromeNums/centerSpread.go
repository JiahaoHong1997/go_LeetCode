package palindromeNums

func countSubstringsCS(s string) int {
	n := len(s)
	ret := 0
	for i:=0; i<n; i++ {
		ret += expandAroundCenter(s,i,i)
		ret += expandAroundCenter(s,i,i+1)
	}
	return ret
}

func expandAroundCenter(s string, left int, right int) int {
	count := 0
	for ; left >= 0 && right < len(s) && s[left] == s[right]; {
		left--
		right++
		count++
	}
	return count
}