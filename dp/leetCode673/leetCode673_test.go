package leetcode673

import (
	"reflect"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	in := [][]int{
		{1,3,5,4,7},
		{0,1,0,3,2,3},
		{7,7,7,7,7,7,7},
	}

	got := []int{}
	want := []int{2,1,7}

	for i:=0; i<len(in); i++ {
		got = append(got, findNumberOfLIS(in[i]))
	}

	if !reflect.DeepEqual(want,got) {
		t.Errorf("expected%v, got%v", want, got)
	}
}
