package ci17

import "math"

func printNumbers(n int) []int {
	ret := []int{}

	x := math.Pow10(n)
	for i:=1; i<x; i++ {
		ret = append(ret, i)
	}
	return ret
}