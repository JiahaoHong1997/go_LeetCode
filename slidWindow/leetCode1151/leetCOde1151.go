package leetcode1151

func minSwaps(data []int) int {

	totalOnes := 0
	for i:=0; i<len(data); i++ {
		if data[i] == 1 {
			totalOnes++
		}
	}
	
	l, r := 0, totalOnes
	countOnes := 0
	for i:=l; i<r; i++ {
		if data[i] == 1 {
			countOnes++
		}
	}

	maxWindowOnes := countOnes

	for r < len(data) {
		x := data[r]
		r++

		countOnes += x-data[l]
		maxWindowOnes = max(maxWindowOnes, countOnes)

		l++
	}

	return totalOnes-maxWindowOnes
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}