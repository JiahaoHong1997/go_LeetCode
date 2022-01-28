package ci05

func replaceSpace(s string) string {
	r := []byte{'%','2','0'}
	sByte := []byte(s)

	for i:=0; i<len(sByte); i++ {
		x := sByte[i]
		if x == ' ' {
			sByte = append(sByte[:i], sByte[i+1:]...)
			sByte = append(sByte[:i], append(r, sByte[i:]...)...)
		}
	}
	return string(sByte)
}