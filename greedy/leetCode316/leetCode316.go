package leetcode316

func removeDuplicateLetters(s string) string {
	count := [26]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}

	stack := []byte{}
	instack := [26]bool{}
	for i:=0; i<len(s); i++ {
		ch := s[i]
		if !instack[ch-'a'] {
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				x := stack[len(stack)-1] - 'a'
				if count[x] == 0 {
					break
				}
				stack = stack[:len(stack)-1]
				instack[x] = false
			}
			stack = append(stack, ch)
			instack[ch-'a'] = true
		}
		count[ch-'a']--
	}
	return string(stack)
}
