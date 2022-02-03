package ci29

func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []int{}
    }
	var (
		top = 0
		bottom = len(matrix) - 1
		left = 0
		right = len(matrix[0]) - 1
	)


	ret := []int{}
	isAsc1 := true
	isAsc2 := true
	for top <= bottom && left <= right {
		if top == bottom {
			if isAsc1 {
				for left <= right {
					ret = append(ret, matrix[top][left])
					left++
				}
			} else {
				for left <= right {
					ret = append(ret, matrix[top][right])
					right--
				}
			}
			break
		}
		if left == right {
			if isAsc2 {
				for top <= bottom {
					ret = append(ret, matrix[top][left])
					top++
				}
			} else {
				for top <= bottom {
					ret = append(ret, matrix[bottom][left])
					bottom--
				}
			}
			break
		}

		for i:=left; i<right; i++ {
			ret = append(ret, matrix[top][i])
		}
		isAsc1 = false

		for i:=top; i<bottom; i++ {
			ret = append(ret, matrix[i][right])
		}
		isAsc2 = false

		for i:=right; i>left; i-- {
			ret = append(ret, matrix[bottom][i])
		}
		isAsc1 = true

		for i:=bottom; i>top; i-- {
			ret = append(ret, matrix[i][left])
		}
		isAsc2 = true

		top++
		bottom--
		left++
		right--
	}
	return ret
}