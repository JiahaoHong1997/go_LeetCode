package ci58


func reverseLeftWords(s string, n int) string {
	b := []byte(s)

	var reverse func([]byte) 
	reverse = func (bb []byte)  {
		for i:=0; i<len(bb)/2; i++ {
			bb[i], bb[len(bb)-i-1] = bb[len(bb)-i-1], bb[i]
		}
	}

	reverse(b[:n])
	reverse(b[n:])
	reverse(b)
	return string(b)
}