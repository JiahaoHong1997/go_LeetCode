package leetcode1388

func maxSizeSlices(slices []int) int {
	dp0 := make([][]int, len(slices))
	dp1 := make([][]int, len(slices))

	choose := len(slices)/3
	for i:=0; i<len(dp0); i++ {
		dp0[i] = make([]int, choose+1)
		dp1[i] = make([]int, choose+1)
	}

	dp0[0][1] = slices[0]
	dp1[1][1] = slices[1]
	
	for i:=1; i<len(slices); i++ {
		for j:=1; j<=choose && j <= i ; j++ {
			if i-2 < 0 {
				dp0[i][j] = dp0[i-1][j]
			} else if i < len(slices)-1 {
				dp0[i][j] = max(dp0[i-1][j], dp0[i-2][j-1]+slices[i])
			}
			
			if i>1 && i<3 {
				dp1[i][j] = dp1[i-1][j]
			} else if i >= 3 {
				dp1[i][j] = max(dp1[i-1][j], dp1[i-2][j-1]+slices[i])
			}
		}
	}

	return max(dp0[len(dp0)-2][choose],dp1[len(dp1)-1][choose])
}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}