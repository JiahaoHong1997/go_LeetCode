package leetcode1901

func findPeakGrid(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	var rowMax func(int) int
	rowMax = func(i int) int {
		if i < 0 || i >= m {
			return -1
		} 
		
		c := 0
		for j:=1; j<n; j++ {
			if mat[i][j] > mat[i][c] {
				c = j
			}
		}
		return c
	}

	l, r := 0, m-1
	for l <= r {
		mid := l + (r-l)/2
		up, cur, down := rowMax(mid-1), rowMax(mid), rowMax(mid+1)
		upVal, curVal, downVal := -1, -1, -1
		if up != -1 {
			upVal = mat[mid-1][up]
		}
		if cur != -1 {
			curVal = mat[mid][cur]
		}
		if down != -1 {
			downVal = mat[mid+1][down]
		}
		if curVal >= upVal && curVal >= downVal {
			return []int{mid,cur}
		}

		if upVal > curVal && upVal > downVal {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return []int{}
}