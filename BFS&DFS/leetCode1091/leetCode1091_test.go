package leetCode1091

import (
	"testing"
)

func TestSPBM(t *testing.T) {
	in := [][]int{
		{0,0,0},{1,0,0},{1,1,0},
	}

	got := shortestPathBinaryMatrix(in)
	want := 3

	if got != want {
		t.Errorf("expected:%v, got%v", want, got)
	}
}

func benchmark(b *testing.B, grid [][]int) {
	for i:=0; i<b.N; i++ {
		shortestPathBinaryMatrix(grid)
	}
}

func benchmarkMap(b *testing.B, grid [][]int) {
	for i:=0; i<b.N; i++ {
		shortestPathBinaryMatrixMap(grid)
	}
}

func BenchmarkSPBMwithoutMap(b *testing.B) {
	in := [][]int{
		{0,0,0},{1,0,0},{1,1,0},
	}
	benchmark(b,in)
}

func BenchmarkSPBMwithMap(b *testing.B) {
	in := [][]int{
		{0,0,0},{1,0,0},{1,1,0},
	}
	benchmarkMap(b,in)
}