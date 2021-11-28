package leetcode40

import (
	"reflect"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	in := []int{10,1,2,7,6,1,5}
	target := 8
	want := [][]int{{1,1,6},
		{1,2,5},
		{1,7},
		{2,6}}

	got := combinationSum2(in, target)
	if !reflect.DeepEqual(want,got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}