package palindromeNums

import (
	"testing"
)


func benchmarkDP(b *testing.B, s string) {
	for i:=0; i<b.N; i++ {
		countSubstringsDP(s)
	}
}

func benchmarkCS(b *testing.B, s string) {
	for i:=0; i<b.N; i++ {
		countSubstringsCS(s)
	}
}

func BenchmarkCountSubstringsDP(b *testing.B) {
	in := []string{
		"abc",
		"aaa",
		"babad",
		"cbbd",
		"shduvbsdfuudsxneuuenx",
	}
	for i:=0; i<len(in); i++ {
		benchmarkDP(b,in[i])
	}
	
}

func BenchmarkCountSubstringsCS(b *testing.B) {
	in := []string{
		"abc",
		"aaa",
		"babad",
		"cbbd",
		"shduvbsdfuudsxneuuenx",
	}
	for i:=0; i<len(in); i++ {
		benchmarkCS(b,in[i])
	}	
}