package leetcode254

import (
	"reflect"
	"testing"
)

func TestGetFactors(t *testing.T) {
	n := []int{1,37,12,32}
	want := [][][]int{
		{},
		{},
		{{2,6},{2,2,3},{3,4}},
		{{2,16},{2,2,8},{2,2,2,4},{2,2,2,2,2},{2,4,4},{4,8}},
	}
	got := [][][]int{}
	for i:=0; i<len(n); i++ {
		tmp := getFactors(n[i])
		got = append(got, tmp)
	}

	if !reflect.DeepEqual(want,got) {
		t.Errorf("expected:%v, got:%v", want,got)
	}
}