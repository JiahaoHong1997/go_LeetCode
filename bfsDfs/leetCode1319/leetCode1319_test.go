package leetcode1319

import (
	"reflect"
	"testing"
)

func TestMakeConnected(t *testing.T) {
	n := []int{
		4,
		6,
		6,
		5,
		11,
	}

	connections := [][][]int{
		{{0,1},{0,2},{1,2}},
		{{0,1},{0,2},{0,3},{1,2},{1,3}},
		{{0,1},{0,2},{0,3},{1,2}},
		{{0,1},{0,2},{3,4},{2,3}},
		{{1,4},{0,3},{1,3},{3,7},{2,7},{0,1},{2,4},{3,6},{5,6},{6,7},{4,7},{0,7},{5,7}},
	}

	want := []int{
		1,
		2,
		-1,
		0,
		3,
	}

	got := []int{}

	for i:=0; i<len(n); i++ {
		got = append(got, makeConnected(n[i],connections[i]))
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}