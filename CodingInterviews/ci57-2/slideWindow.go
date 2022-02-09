package ci57

func findContinuousSequenceTP(target int) [][]int {
	l, r := 1, 2

	ret := [][]int{}
	for l < r {
		sum := (l+r)*(r-l+1)/2
		if sum == target {
			t := []int{}
			for i:=l; i<=r; i++ {
				t = append(t, i)
			}
			ret = append(ret, t)
            l++
		} else if sum < target {
			r++
		} else {
			l++
		}
	}
	return ret
}