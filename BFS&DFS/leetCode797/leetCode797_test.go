package leetCode797

import (
	"testing"
	"reflect"
)

func TestAllPathsSourceTarget(t *testing.T) {
	in1 := [][]int{{1,2},{3},{3},{}}
	in2 := [][]int{{4,3,1},{3,2,4},{3},{4},{}}
	in3 := [][]int{{1},{}}
	in4 := [][]int{{1,2,3},{2},{3},{}}
	in5 := [][]int{{1,3},{2},{3},{}}

	want1 := [][]int{{0,1,3},{0,2,3}}
	want2 := [][]int{{0,4},{0,3,4},{0,1,3,4},{0,1,2,3,4},{0,1,4}}
	want3 := [][]int{{0,1}}
	want4 := [][]int{{0,1,2,3},{0,2,3},{0,3}}
	want5 := [][]int{{0,1,2,3},{0,3}}

	got1 := allPathsSourceTarget(in1)
	got2 := allPathsSourceTarget(in2)
	got3 := allPathsSourceTarget(in3)
	got4 := allPathsSourceTarget(in4)
	got5 := allPathsSourceTarget(in5)

	if !reflect.DeepEqual(got1, want1) {
		t.Errorf("expected:%v, got:%v", want1, got1)
	}
	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("expected:%v, got:%v", want2, got2)
	}
	if !reflect.DeepEqual(got3, want3) {
		t.Errorf("expected:%v, got:%v", want3, got3)
	}
	if !reflect.DeepEqual(got4, want4) {
		t.Errorf("expected:%v, got:%v", want4, got4)
	}
	if !reflect.DeepEqual(got5, want5) {
		t.Errorf("expected:%v, got:%v", want5, got5)
	}
}

func benchmark(b *testing.B, in [][]int) {
	for i:=0; i<b.N; i++ {
		allPathsSourceTarget(in)
	}
}

func Benchmark1(b *testing.B) {
	in1 := [][]int{{1,2},{3},{3},{}}
	benchmark(b,in1)
}

func Benchmark2(b *testing.B) {
	in2 := [][]int{{4,3,1},{3,2,4},{3},{4},{}}
	benchmark(b,in2)
}

func Benchmark3(b *testing.B) {
	in3 := [][]int{{1},{}}
	benchmark(b,in3)
}

func Benchmark4(b *testing.B) {
	in4 := [][]int{{1,2,3},{2},{3},{}}
	benchmark(b,in4)
}

func Benchmark5(b *testing.B) {
	in5 := [][]int{{1,3},{2},{3},{}}
	benchmark(b,in5)
}

