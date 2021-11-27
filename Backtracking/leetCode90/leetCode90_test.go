package leetCode90

import (
	"reflect"
	"testing"
)


func TestSubsetsWithDup(t *testing.T) {
	in := []int{4,4,4,1,4}
	got := subsetsWithDup(in)
	want := [][]int{{},{1},{1,4},{1,4,4},{1,4,4,4},{1,4,4,4,4},{4},{4,4},{4,4,4},{4,4,4,4}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected:%v,  got:%v", want, got)
	}
}