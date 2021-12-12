package leetcode1231

import (
	"reflect"
	"testing"
)

func TestMaximizeSweetness(t *testing.T) {
	s := [][]int{
		{1,2,3,4,5,6,7,8,9},
		{5,6,7,8,9,1,2,3,4},
		{1,2,2,1,2,2,1,2,2},
	}

	k := []int{5,8,2}
	want := []int{6,1,5}
	got := []int{} 

	for i:=0; i<len(k); i++ {
		got = append(got, maximizeSweetness(s[i],k[i]))
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}