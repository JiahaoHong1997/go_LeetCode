package ci56

import "math"

func singleNumber(nums []int) int {
	count := [32]int{}
	for i:=0; i<len(nums); i++ {
		t := nums[i]
		j := 0
		for t > 0 {
			if t^(t-1) == 1 {
				count[j]++
			} 
			j++
			t >>= 1
		}
	}

	ret := 0
	for i:=0; i<len(count); i++ {
		x := count[i] % 3
		ret += x * int(math.Pow(2,float64(i)))
	}
	return ret
}