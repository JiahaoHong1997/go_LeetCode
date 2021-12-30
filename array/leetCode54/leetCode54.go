package leetcode54

func spiralOrder(matrix [][]int) []int {
	var (
		high = 0
		low = len(matrix)-1
		left = 0
		right = len(matrix[0])-1
	)

	ret := make([]int,0)
	down := true
	turnRight := true
	for low >= high && right >= left {
		if low == high {
			var x, y int
			if turnRight {
				x, y = left, right
			} else {
				x, y = right, left
			}
			for i := x; i <= y; i++ {
				ret = append(ret, matrix[low][i])
			}
			break
		}

		if left == right {
			var x, y int
			if down {
				x, y = high, low
			} else {
				x, y = low, high
			}
			for i:=x; i<=y; i++ {
				ret = append(ret, matrix[i][left])
			}
			break
		}

		for i := left; i<right; i++ {
			ret = append(ret, matrix[high][i])
			turnRight = false
		}
		

		for i:=high; i<low; i++ {
			ret = append(ret, matrix[i][right])
			down = false

		}
		
		for i:=right; i>left; i-- {
			ret = append(ret, matrix[low][i])
			turnRight = true
		}


		for i:=low; i>high; i-- {
			ret = append(ret, matrix[i][left])
			down = true
		}
		high++
		left++
		low--
		right--
	} 
	 
	return ret
}