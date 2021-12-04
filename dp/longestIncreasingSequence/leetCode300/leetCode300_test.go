package leetcode300

import (
	"reflect"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	in := [][]int{
		{10,9,2,5,3,7,101,18},
		{0,1,0,3,2,3},
		{7,7,7,7,7,7,7},
	}

	got := []int{}
	want := []int{4,4,1}

	for i:=0; i<len(in); i++ {
		got = append(got, lengthOfLIS(in[i]))
	}

	if !reflect.DeepEqual(want,got) {
		t.Errorf("expected%v, got%v", want, got)
	}
}

func benchmark1(b *testing.B, in [][]int) {
	for i:=0; i<b.N; i++ {
		for j:=0; j<len(in); j++ {
			lengthOfLISDP(in[j])
		}
	}
}

func benchmark2(b *testing.B, in [][]int) {
	for i:=0; i<b.N; i++ {
		for j:=0; j<len(in); j++ {
			lengthOfLIS(in[j])
		}
	}
}

func Benchmark1(b *testing.B) {
	in := [][]int{
		{10,9,2,5,3,7,101,18},
		{0,1,0,3,2,3},
		{7,7,7,7,7,7,7},
		{3,6,12,21,15,13,14,12,24,17,22,26,5,7,23,135,631,1,212,45,9,24,543,27,46,27,42,53,44,63,23,74,34,2,78,42,14,56,32,45,63,73,34,74},
	}
	benchmark1(b,in)
}

func Benchmark2(b *testing.B) {
	in := [][]int{
		{10,9,2,5,3,7,101,18},
		{0,1,0,3,2,3},
		{7,7,7,7,7,7,7},
		{3,6,12,21,15,13,14,12,24,17,22,26,5,7,23,135,631,1,212,45,9,24,543,27,46,27,42,53,44,63,23,74,34,2,78,42,14,56,32,45,63,73,34,74},
	}
	benchmark2(b,in)
}