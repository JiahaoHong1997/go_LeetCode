package ci50

func firstUniqChar(s string) byte {

	if s == "" {
		return ' '
	}
	m := [26]int{}
	for i:=0; i<len(s); i++ {
		m[s[i]-'a']++
	}

	for i:=0; i<len(s); i++ {
		if m[s[i]-'a'] == 1 {
			return s[i]
		}
	}

	return ' '
}