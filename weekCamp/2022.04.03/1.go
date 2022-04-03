package weekCamp

import (
	"strings"
	"strconv"
	"math"
)

func convertTime(current string, correct string) int {
    s1 := strings.Split(current, ":")
    h1, m1 := s1[0], s1[1]
    
    s2 := strings.Split(correct, ":")
	h2, m2 := s2[0], s2[1]
	
	nowh, _ := strconv.Atoi(h1)
	nowm, _ := strconv.Atoi(m1)
	now := nowh*60+nowm

	th, _ := strconv.Atoi(h2)
	tm, _ := strconv.Atoi(m2)
	t := th*60+tm

	diff := abs(now, t)
	choose := []int{1,5,15,60}
	
	dp := make([]int, diff+1)
	dp[0] = 0

	for i:=1; i<=diff; i++ {
		dp[i] = math.MaxInt32
		for j:=len(choose)-1; j>=0; j-- {
			if choose[j] > i {
				continue
			}
			dp[i] = min(dp[i], dp[i-choose[j]]+1)
		}
		
	}

	return dp[diff]
}

func abs(a, b int) int {
	if a > b {
		return a-b 
	} 
	return b-a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}