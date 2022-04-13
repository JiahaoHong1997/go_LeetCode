package leetcode186

func reverseWords(s []byte) {

	var reverse func([]byte)
	reverse = func(b []byte) {
		for i:=0; i<len(b)/2; i++ {
			b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
		}
	}

	first := 0
	for i:=0; i<=len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			reverse(s[first:i])
			first = i+1
		}
	}

	reverse(s)
}