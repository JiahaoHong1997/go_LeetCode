package weekCamp

import "sort"

func countRectangles(rectangles [][]int, points [][]int) []int {

	length := len(rectangles)
	m := make([][]int, 101)
	sorted := make([]bool, 101)

	for i:=0; i<length; i++ {
		x, y := rectangles[i][0], rectangles[i][1]
		for j:=0; j<=y; j++ {
			m[j] = append(m[j], x)
		}
	}

	var judgeCount func([]int, int) int
	judgeCount = func(arr []int, target int) int {
		r := len(arr)-1
		if target > arr[r] {
			return 0
		}

		n := sort.Search(len(arr), func(i int) bool {
			return arr[i] >= target
		})

		if n == len(arr) {
			return 0
		}
		
		return len(arr)-n
	}



	ret := make([]int, len(points))

	for i:=0; i<len(points); i++ {
		x, y := points[i][0], points[i][1]
		if len(m[y]) == 0 {
			continue
		}

		if !sorted[y] {
			sort.Ints(m[y])
			sorted[y] = true
		}

		ret[i] = judgeCount(m[y], x)
	}

	return ret
}