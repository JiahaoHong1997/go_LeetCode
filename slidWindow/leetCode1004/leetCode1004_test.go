package leetcode1004

import "testing"

func TestLongestOnes(t *testing.T) {
	nums := []int{1,1,0,1,0,1,1,1,0,0,1,0,1,1,1,1,1,0,1,0,0,0,1,1,0,1,0,1,1,0,1}
	k := 3

	want := 8
	got := longestOnes(nums,k)
	if want != got {
		t.Fatalf("expected:%v, got:%v", want, got)
	}
}