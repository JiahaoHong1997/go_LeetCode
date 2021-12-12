package leetcode1182

import (
	"reflect"
	"testing"
)

func TestShortestDistanceColor(t *testing.T) {
	colors := [][]int{
		{1,1,2,1,3,2,2,3,3},
		{1,2},
		{2,1,2,2,1},
	}
	queries := [][][]int{
		{{1,3},{2,2},{6,1}},
		{{0,3}},
		{{1,1},{4,3},{1,3},{4,2},{2,1}},
	} 
	want := [][]int{
		{3,0,3},
		{-1},
		{0,-1,-1,1,1},
	}
	got := make([][]int,len(colors))

	for i:=0; i<len(colors); i++ {
		got[i] = append(got[i], shortestDistanceColor(colors[i],queries[i])...)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}