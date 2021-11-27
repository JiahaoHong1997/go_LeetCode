package leetCode90

import (
	"reflect"
	"testing"
)


func TestSubsetsWithDup(t *testing.T) {
	in := []int{1,2,2}
	got := subsetsWithDup(in)
	want := [][]int{{},{1},{1,2},{1,2,2},{2},{2,2}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected:%v,  got:%v", want, got)
	}
}