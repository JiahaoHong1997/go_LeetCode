package leetcode679

func judgePoint24(cards []int) bool {


	set := make([]float64, 4)
	for i:=0; i<4; i++ {
		set[i] = float64(cards[i])
	}
	function := []byte{'+','-','*','/'}


	var backTracking func() bool
	backTracking = func() bool {

		if len(set) == 1 {
			if set[0] - 24.0 <= 1e-6 && set[0] - 24.0 >= -1*1e-6 {
				return true
			}
			return false
		}

		for i:=0; i<len(set); i++ {
			x := set[i]
			set = append(set[:i], set[i+1:]...)
			for j:=0; j<len(set); j++ {
				y := set[j]
				set = append(set[:j], set[j+1:]...)
				for k:=0; k<4; k++ {
					if k == 3 && y <= 1e-6 {
						continue
					}
					newNum := cal(function[k], x, y)
					set = append(set, newNum)
					if backTracking() {
						return true
					}
					set = set[:len(set)-1]
				}
				set = append(set[:j], append([]float64{y}, set[j:]...)...)
			}
			set = append(set[:i], append([]float64{x}, set[i:]...)...)
		}
		return false
	}

	return backTracking()
}

func cal(fun byte, a,b float64) float64 {
	switch {
	case fun == '+': return a+b
	case fun == '-': return a-b
	case fun == '*': return a*b
	case fun == '/': return a/b 
	}
	return 0
}