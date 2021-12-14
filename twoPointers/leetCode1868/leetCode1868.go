package leetcode1868



func findRLEArray(encoded1 [][]int, encoded2 [][]int) [][]int {
	n1, n2 := len(encoded1), len(encoded2)
	ret := [][]int{}

	for i,j := 0, 0; i < n1 || j < n2; {
		n := len(ret)

		if i == n1 {
			x1, x2 := encoded1[i-1][0], encoded2[j][0]
			y1, y2 := encoded1[i-1][1], encoded2[j][1]
			if x1*x2 == ret[n-1][0] {
				ret[n-1][1] += y2 
			} else {
				ret = append(ret, []int{x1*x2,y2})
			}
			if y1 == y2 {
				break
			} else {
				encoded1[i-1][1] -= y2
				j++
				continue
			}
		} else if j == n2 {
			x1, x2 := encoded1[i][0], encoded2[j-1][0]
			y1, y2 := encoded1[i][1], encoded2[j-1][1]
			if x1*x2 == ret[n-1][0] {
				ret[n-1][1] += y1 
			} else {
				ret = append(ret, []int{x1*x2,y1})
			}
			if y1 == y2 {
				break
			} else {
				encoded2[j-1][1] -= y1
				i++
				continue
			}			
		}

		if i == 0 && j == 0 {
			ret = append(ret, []int{encoded1[0][0]*encoded2[0][0], min(encoded1[0][1],encoded2[0][1])})
			if encoded1[0][1] < encoded2[0][1] {
				i++
				encoded2[0][1] -= encoded1[0][1]
			} else if encoded1[0][1] > encoded2[0][1] {
				j++
				encoded1[0][1] -= encoded2[0][1]
			} else {
				i++
				j++
			}
			continue
		}

		
		x1, x2 := encoded1[i][0], encoded2[j][0]
		y1, y2 := encoded1[i][1], encoded2[j][1]
		if x1*x2 == ret[n-1][0] {
			ret[n-1][1] += min(y1,y2)
		} else {
			ret = append(ret, []int{x1*x2, min(y1,y2)})
		}

		if y1 < y2 {
			encoded2[j][1] -= y1
			i++
		} else if y1 > y2 {
			encoded1[i][1] -= y2
			j++
		} else {
			i++
			j++
		}
	}
    
	return ret
}

func min(a,b int) int {
	if a < b{
		return a
	}
	return b
}