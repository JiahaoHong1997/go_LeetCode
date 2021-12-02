package leetcode413

import (
	"testing"
)


func benchmarkNumberOfArithmeticSlices(b *testing.B, nums []int) {
	for i:=0; i<b.N; i++ {
		numberOfArithmeticSlices(nums)
	}
}

func benchmarkNumberOfArithmeticSlicesBetter(b *testing.B, nums []int) {
	for i:=0; i<b.N; i++ {
		numberOfArithmeticSlicesBetter(nums)
	}
}

func BenchmarkCountSubstringsDP(b *testing.B) {
	in := [][]int{
		{1,2,3,4},
		{1},
		{1,3,5,7,9,11,2,3,6,2,3,8,10,12,44,12,20,28,36},
		{1},
	}
	for i:=0; i<len(in); i++ {
		benchmarkNumberOfArithmeticSlices(b,in[i])
	}
	
}

func BenchmarkCountSubstringsCS(b *testing.B) {
	in := [][]int{
		{1,2,3,4},
		{1},
		{1,3,5,7,9,11,2,3,6,2,3,8,10,12,44,12,20,28,36},
		{1},
	}
	for i:=0; i<len(in); i++ {
		benchmarkNumberOfArithmeticSlicesBetter(b,in[i])
	}	
}