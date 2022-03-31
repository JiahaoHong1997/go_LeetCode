package leetcode498

func findDiagonalOrder(mat [][]int) []int {

	ret := []int{}
	m, n := len(mat), len(mat[0])
	if m == 1 {
		return mat[0]
	} else if n == 1 {
		for i:=0; i<m; i++ {
			ret = append(ret, mat[i][0])
		}
		return ret
	}

	i, j := 0, 0
	direction := true
	rowHalf := true
	columnHalf := true

	for i<m && j<n {
		ret = append(ret, mat[i][j])
		if (i == 0 && direction && rowHalf) || (j == n-1 && direction && !rowHalf) {
			if rowHalf {
				if j < n-1 {
					j++
				} else {
					i++
				}
			} else {
				i++
			}
			direction = false
		} else if (j == 0 && !direction) || (i == m-1 && !direction && !columnHalf) {
			if columnHalf {
				if i < m-1 {
					i++
				} else {
					j++
				}
			} else {
				j++
			}
			direction = true
		} else if direction {
			i--
			j++
		} else {
			i++
			j--
		}
		
		if i+j >= m-1 {
			columnHalf = false
		}
		if i+j >= n-1 {
			rowHalf = false
		}
	}
	
	return ret
}