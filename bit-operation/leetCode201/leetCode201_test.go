package leetcode201

import "testing"

func benchmark1(b *testing.B, left,right []int) {
	for i:=0; i<b.N; i++ {
		for k, v := range left {
			rangeBitwiseAnd(v, right[k])
		} 
	}
}

func benchmark2(b *testing.B, left,right []int) {
	for i:=0; i<b.N; i++ {
		for k, v := range left {
			rangeBitwiseAndOP(v, right[k])
		} 
	}
}

var left []int
var right []int

func init() {
	left = []int{5,0,1,2,3}
	right = []int{4325324634,4327854637128,43267452378,6785236523896,589205732980572}
}

func Benchmark1(b *testing.B) {
	benchmark1(b, left, right)
}

func Benchmark2(b *testing.B) {
	benchmark2(b, left, right)
}