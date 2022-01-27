package leetcode494

import "testing"

var (
	nums   = []int{15, 42, 44, 50, 34, 46, 10, 27, 27, 27, 27, 40, 3, 35, 4, 47, 32, 45, 13, 46}
	target = 26
)

func BenchmarkDP(b *testing.B) {
	benchmarkDP(b)
}

func BenchmarkBT(b *testing.B) {
	benchmarkBT(b)
}

func benchmarkDP(b *testing.B) {
	for n := 0; n < b.N; n++ {
		findTargetSumWaysDP(nums, target)
	}
}

func benchmarkBT(b *testing.B) {
	for n := 0; n < b.N; n++ {
		findTargetSumWaysBT(nums, target)
	}
}
