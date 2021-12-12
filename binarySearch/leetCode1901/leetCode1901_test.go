package leetcode1901

import (
	"reflect"
	"testing"
)

func TestFindPeakGrid(t *testing.T) {
	mat := [][]int{
		{10,20,15},{21,30,14},{7,16,32},
	}

	want := []int{2,2}

	got := findPeakGrid(mat)

	if !reflect.DeepEqual(want,got) {
		t.Errorf("expected:%v,got:%v", want, got)
	}
}