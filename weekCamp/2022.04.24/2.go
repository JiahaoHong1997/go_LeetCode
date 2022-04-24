package weekCamp

import "sort"

func countLatticePoints(circles [][]int) (ans int) {

	// 按半径从大到小排序，这样能更早遇到包含 (x,y) 的圆
	sort.Slice(circles, func(i, j int) bool { return circles[i][2] > circles[j][2] })

	for x := 0; x <= 200; x++ {
		for y := 0; y <= 200; y++ {
			for _, c := range circles {
				// 判断 (x,y) 是否在圆 c 内
				if (x-c[0])*(x-c[0])+(y-c[1])*(y-c[1]) <= c[2]*c[2] {
					ans++
					break
				}
			}
		}
	}

	return
}