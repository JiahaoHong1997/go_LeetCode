package leetcode39

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	in := []int{2,3,6,7}
	target := 7
	want := [][]int{{2,2,3},{7}}

	got := combinationSum(in, target)
	if !reflect.DeepEqual(want,got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}