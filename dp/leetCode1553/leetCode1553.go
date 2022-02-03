package leetcode1553


func minDays(n int) int {
	
	m := make(map[int]int)
	m[0], m[1] = 0, 1
	var backTracking func(int) int
	backTracking = func(num int) int{
		if d, ok := m[num]; ok {
			return d
		} 

		m[num] = min(backTracking(num/2)+num%2+1, backTracking(num/3)+num%3+1) 
		return m[num]
	}
	
	return backTracking(n)
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}