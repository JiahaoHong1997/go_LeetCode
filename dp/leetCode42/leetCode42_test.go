package leetcode42

import (
	"reflect"
	"testing"
)

func TestTrap(t *testing.T) {
	height := [][]int{
		{0,1,0,2,1,0,1,3,2,1,2,1},
		{4,2,0,3,2,5},
	}
	want := []int{6,9}
	got := []int{}
	for i:=0; i<len(height); i++ {
		got =append(got, trap(height[i]))
	}

	if !reflect.DeepEqual(want,got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}