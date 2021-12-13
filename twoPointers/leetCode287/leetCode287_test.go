package leetcode287

import (
	"reflect"
	"testing"
)
func TestFindDuplicate(t *testing.T) {
	nums := [][]int{
		{1,3,4,2,2},
		{1,3,4,2,4},
		{3,1,3,4,2},
		{1,1},
		{2,2,2,2,2,2,2,2,2,2},
		{1,1,2},
	}
	want := []int{2,4,3,1,2,1}

	got := []int{}
	for i:=0; i<len(nums); i++ {
		got = append(got, findDuplicate(nums[i]))
	}

	if !reflect.DeepEqual(want,got) {
		t.Errorf("expected:%v,got:%v",want,got)
	}
}