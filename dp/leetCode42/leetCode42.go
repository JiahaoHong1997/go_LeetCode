package leetcode42

func trap(height []int) int {
    ldp := make([]int,len(height))
    rdp := make([]int,len(height))

	ldp[0] = height[0]
	rdp[len(height)-1] = height[len(height)-1]

	for i:=1; i<len(height); i++ {
		ldp[i] = max(ldp[i-1],height[i])
		rdp[len(height)-1-i] = max(rdp[len(height)-i],height[len(height)-1-i])
	}

	ret := 0
	for i := 0; i<len(height); i++ {
		ret += min(ldp[i],rdp[i]) - height[i]
	}
	return ret
}

func max(a,b int) int {
	if a > b {
		return a
	} 

	return b
}

func min(a,b int) int {
	if a< b {
		return a
	}
	return b
}