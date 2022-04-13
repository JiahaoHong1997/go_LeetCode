package leetcode256

func minCost(costs [][]int) int {

	a, b, c := costs[0][0], costs[0][1], costs[0][2]
	for i:=1; i<len(costs); i++ {
		a, b, c = min(b, c)+costs[i][0], min(a, c)+costs[i][1], min(a, b)+costs[i][2]
	}

	return min(a, min(b, c))
}

  

func min(a, b int) int {
	if a < b {
		return a
	}
	
	return b
}