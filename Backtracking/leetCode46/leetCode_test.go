package leetcode46

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	in := []int{1,2,3}
	want := [][]int{{1,2,3},{1,3,2},{2,1,3},{2,3,1},{3,1,2},{3,2,1}}

	got := permute(in)
	if !reflect.DeepEqual(want,got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}