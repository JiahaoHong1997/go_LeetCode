package leetcode1388

import "testing"

func TestMaxSizeSlices(t *testing.T) {
	in := []int{3,7,8,6,3,10,4,2,9}
	want := 27
	got := maxSizeSlices(in)
	if want != got {
		t.Errorf("expected:%v, got:%v",want, got)
	}
}