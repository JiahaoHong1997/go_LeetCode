package leetcode78

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	in := []int{1, 2, 3}
	got := subsets(in)
	want := [][]int{{},{1},{1,2},{1,2,3},{1,3},{2},{2,3},{3}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected:%v,  got:%v", want, got)
	}
}
