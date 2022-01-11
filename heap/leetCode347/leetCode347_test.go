package leetcode347

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	nums := [][]int{
		{1,1,1,2,2,3},
		{1},
	}

	k := []int{
		2,1,
	}

	want := [][]int{
		{2,1},
		{1},
	}

	got := [][]int{}

	for i:=0; i<len(k); i++ {
		got = append(got, topKFrequent(nums[i], k[i]))
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}