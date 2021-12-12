package leetcode1060

import (
	"reflect"
	"testing"
)

func TestMissingElement(t *testing.T) {
	nums := [][]int{
		{4,7,9,10},
		{4,7,9,10},
		{1,2,4},
	}

	k := []int{1,3,3}
	want := []int{5,8,6}
	got := []int{}

	for i:=0; i<len(nums); i++ {
		got = append(got, missingElement(nums[i], k[i]))
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v",want, got)
	}

}