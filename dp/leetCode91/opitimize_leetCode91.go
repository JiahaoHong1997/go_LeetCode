package leetcode91

func numDecodingsOP(s string) int {
    if s[0] == '0' {
		return 0
	}

	dp0 := 1
	dp1 := 0

	for i := 1; i < len(s); i++ {
		if s[i] == '0' && (s[i-1] == '0' || int(s[i-1])-'2' > 0) { // 00 || 30
			return 0
		}

		if s[i] == '0' {	
			dp0,dp1 = 0,dp0
		} else {
			if s[i-1] == '0' || int(s[i-1])-'2' > 0 || (int(s[i-1])-'2' == 0 && int(s[i])-'6' >= 1) {  // 0a || 3a || (27~29)
				dp0,dp1 = dp0+dp1, 0
			} else {
				dp0,dp1 = dp0+dp1, dp0
			}
		}
	}
	return dp0+dp1
}