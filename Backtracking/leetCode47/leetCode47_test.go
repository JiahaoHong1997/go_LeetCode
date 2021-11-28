package leetcode47

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	in := []int{1,1,2}
	want := [][]int{{1,1,2},{1,2,1},{2,1,1}}

	got := permuteUnique(in)
	if !reflect.DeepEqual(want,got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}