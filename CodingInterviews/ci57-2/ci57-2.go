package ci57

func findContinuousSequence(target int) [][]int {
	ret := [][]int{}
	path := []int{}

	var backTracking func(int,int)
	backTracking = func(sum int, num int) {

		if sum == 0 {
			t := make([]int, len(path))
			copy(t, path)
			ret = append(ret, t)
			return 
		} else if sum < 0 {
			return 
		}

		path = append(path, num)
		backTracking(sum-num, num+1)
	}

	for i:=1; i<=target/2; i++ {
		path = append(path, i)
		backTracking(target-i, i+1)
		path = []int{}
	}

	return ret
}