package leetcode1868

import (
	"reflect"
	"testing"
)

func TestFindRLEArray(t *testing.T) {
	encoded1 := [][]int{{1, 3}, {2, 1}, {3, 2}}
	encoded2 := [][]int{{2, 3}, {3, 3}}

	want := [][]int{{2, 3}, {6, 1}, {9, 2}}
	got := [][]int{}

	got = findRLEArray(encoded1, encoded2)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}
