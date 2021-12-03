package leetcode91

import "testing"

func benchmark1(b *testing.B, s []string) {
	for i:=0; i<b.N; i++ {
		for j:=0; j<len(s); j++ {
			numDecodings(s[j])
		}
	}
} 

func benchmark2(b *testing.B, s []string) {
	for i:=0; i<b.N; i++ {
		for j:=0; j<len(s); j++ {
			numDecodingsOP(s[j])
		}
	}
} 

func Benchmark1(b *testing.B) {
	s := []string{
		"11106",
		"12",
		"226",
		"0",
		"2321532617214312214621201342233121405",
		"27282930",
		"600",
		"2321532617214312214621201344326875321762233121405",
	}
	benchmark1(b,s)
}

func Benchmark2(b *testing.B) {
	s := []string{
		"11106",
		"12",
		"226",
		"0",
		"2321532617214312214621201342233121405",
		"27282930",
		"600",
		"2321532617214312214621201344326875321762233121405",
	}
	benchmark2(b,s)
}